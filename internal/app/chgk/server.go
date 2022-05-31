package chgk

import (
	"context"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk_api_client"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/telegram"
	"strings"
)

type chgkServer struct {
	pb.UnimplementedChgkServiceServer
	tg   *telegram.Client
	chgk *chgk_api_client.Client
	repo Repository
}

func New(config *config.Config, repo Repository) *chgkServer {
	tg := telegram.New(config.ApiKeys.Telegram)
	cl := chgk_api_client.New("https://db.chgk.info/")
	return &chgkServer{
		pb.UnimplementedChgkServiceServer{},
		tg,
		cl,
		repo,
	}
}

func (s chgkServer) SendMessage(ctx context.Context, in *pb.SendMessageReq) (*pb.Message, error) {
	mes, err := s.tg.SendMessage(in.ChatId, in.Text)
	if err != nil {
		return nil, err
	}
	return mes, nil
}

func (s chgkServer) GetRandomQuestion(ctx context.Context, empty *pb.Empty) (*pb.Question, error) {
	q, err := s.chgk.GetRandomQuestion()
	if err != nil {
		return nil, err
	}

	question := &pb.Question{
		Id:       q.Id,
		Question: q.Question,
		Answer:   q.Answer,
		Authors:  q.Authors,
		Comments: q.Comments,
	}
	s.tg.SendMessage(210281851, question.Question)
	return question, nil

}

func filterNewLines(s string) string {
	return strings.Map(func(r rune) rune {
		switch r {
		case 0x000A, 0x000B, 0x000C, 0x000D, 0x0085, 0x2028, 0x2029:
			return -1
		default:
			return r
		}
	}, s)
}
