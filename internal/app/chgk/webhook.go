package chgk

import (
	"context"
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"strings"
	"time"
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
		if err != nil {
			return &pb.Empty{}, err
		}
		_, err = s.tg.SendMessage(update.Message.From.Id, fmt.Sprintf("ID: %d\nВопрос: %s?\nАвтор(ы):%s", q.ID, q.Question, q.Authors))
		if err != nil {
			return &pb.Empty{}, err
		}
		s.cache[update.Message.From.Id] = q
		time.Sleep(30 * time.Second)
		_, err = s.tg.SendMessage(update.Message.From.Id,
			fmt.Sprintf("К сожалению время вышло!\nКомментарии к вопросу: %s\nПравильный ответ:%s", q.Comments, q.Answer))
		delete(s.cache, update.Message.From.Id)
		if err != nil {
			return &pb.Empty{}, err
		}
	default:
		q, ok := s.cache[update.Message.From.Id]
		if ok && q != nil {
			if strings.Contains(strings.ToLower(q.Answer), strings.ToLower(update.Message.Text)) {
				_, err := s.tg.SendMessage(update.Message.From.Id,
					fmt.Sprintf("Верно! Полный ответ: %s\nКомментарии: %s", q.Answer, q.Comments))
				if err != nil {
					return &pb.Empty{}, err
				}
			} else {
				_, err := s.tg.SendMessage(update.Message.From.Id,
					fmt.Sprintf("Неверно! Попробуйте снова!"))
				if err != nil {
					return &pb.Empty{}, err
				}
			}
		} else {
			_, err := s.tg.SendMessage(update.Message.From.Id,
				fmt.Sprintf("Запросите новый вопрос!"))
			if err != nil {
				return &pb.Empty{}, err
			}
		}
		return &pb.Empty{}, nil
	}
	return &pb.Empty{}, nil
}
