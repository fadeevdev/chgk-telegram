package chgk

import (
	"context"
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"strings"
)

func (s *chgkServer) WebHook(ctx context.Context, update *pb.Update) (*pb.Empty, error) {
	fmt.Printf("received a message: %s\n", update)
	if update.Message == nil {
		return &pb.Empty{}, nil
	}
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
		maskedAnswer := maskString(q.Answer)
		_, err = s.tg.SendMessage(update.Message.From.Id, fmt.Sprintf("ID: %d\nQuestion: %s?\nAuthor(s):%s\nAnswer: %s", q.ID, q.Question, q.Authors, maskedAnswer))
		if err != nil {
			return &pb.Empty{}, err
		}
		return &pb.Empty{}, nil
	case "/top":
		pos, err := s.repo.GetTopPosition(ctx, update.Message.From.Id)
		if err != nil {
			return &pb.Empty{}, err
		}
		_, err = s.tg.SendMessage(update.Message.From.Id, fmt.Sprintf("%s, your position is %d in top, answered questions: %d", pos.Username, pos.Position, pos.Questions))
		if err != nil {
			return &pb.Empty{}, err
		}
		return &pb.Empty{}, nil
	default:
		q := s.cache.Get(update.Message.From.Id)
		if q != nil {
			userAnswer := strings.ToLower(update.Message.Text)
			answer := strings.ToLower(q.Answer)
			if userAnswer == answer {
				err := s.repo.AddToTop(ctx, update.Message.From.Id, q.ID)
				if err != nil {
					return &pb.Empty{}, err
				}
				s.cache.Delete(update.Message.From.Id)
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
