package chgk

import (
	"context"
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"strings"
)

func (s *chgkServer) WebHook(ctx context.Context, update *pb.Update) (*pb.Empty, error) {
	fmt.Printf("received a message: %s\n", update)
	switch update.Message.Text {
	case "/start":
		_, err := s.RegisterUser(ctx, update.Message.From)
		if err != nil && err.Error() == ErrUserAlreadyExists {
			_, err := s.tg.SendMessage(update.Message.From.Id, fmt.Sprintf("Welcome back, %s!", update.Message.From.Username))
			if err != nil {
				return &pb.Empty{}, err
			}
			return &pb.Empty{}, nil
		}
		if err != nil {
			return &pb.Empty{}, err
		}
		_, err = s.tg.SendMessage(update.Message.From.Id, fmt.Sprintf("Welcome, %s!", update.Message.From.Username))
		if err != nil {
			return &pb.Empty{}, err
		}
	case "/question":
		q, err := s.chgk.GetRandomQuestion()
		if err != nil {
			return &pb.Empty{}, err
		}
		_, err = s.repo.SaveQuestion(ctx, q)
		s.cache.Put(update.Message.From.Id, q)
		if err != nil {
			return &pb.Empty{}, err
		}
		_, err = s.tg.SendMessage(update.Message.From.Id, fmt.Sprintf("ID: %d\nQuestion: %s?\nAuthor(s):%s", q.ID, q.Question, q.Authors))
		if err != nil {
			return &pb.Empty{}, err
		}
		return &pb.Empty{}, nil
	case "/top":
		pos, err := s.repo.GetTopPosition(ctx, update.Message.From.Id)
		if err != nil {
			return &pb.Empty{}, err
		}
		_, err = s.tg.SendMessage(update.Message.From.Id, fmt.Sprintf("%s, your position is %d in top, answered questions: %d", pos.FirstName, pos.Position, pos.Questions))
		if err != nil {
			return &pb.Empty{}, err
		}
		return &pb.Empty{}, nil
	default:
		q := s.cache.Get(update.Message.From.Id)
		if q != nil {
			userAnswer := strings.ToLower(update.Message.Text)
			answer := strings.Fields(strings.ToLower(q.Answer))
			answered := false
			for _, a := range answer {
				if userAnswer == a {
					answered = true
					break
				}
			}
			if answered {
				err := s.repo.AddToTop(ctx, update.Message.From.Id, q.ID)
				if err != nil {
					return &pb.Empty{}, err
				}
				_, err = s.tg.SendMessage(update.Message.From.Id,
					fmt.Sprintf("Right! Full answer: %s\nComments: %s", q.Answer, q.Comments))
				if err != nil {
					return &pb.Empty{}, err
				}
			} else {
				_, err := s.tg.SendMessage(update.Message.From.Id,
					fmt.Sprintf("Wrong! Try again!"))
				if err != nil {
					return &pb.Empty{}, err
				}
			}
		} else {
			_, err := s.tg.SendMessage(update.Message.From.Id,
				fmt.Sprintf("Query a new question please!"))
			if err != nil {
				return &pb.Empty{}, err
			}
		}
		return &pb.Empty{}, nil
	}
	return &pb.Empty{}, nil
}
