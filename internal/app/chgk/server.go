package chgk

import (
	"context"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk_api_client"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/telegram"
)

type chgkServer struct {
	pb.UnimplementedChgkServiceServer
	tg    *telegram.Client
	chgk  *chgk_api_client.Client
	repo  Repository
	cache map[uint64]*chgk_api_client.Question
}

func New(config *config.Config, repo Repository) *chgkServer {
	tg := telegram.New(config.ApiKeys.Telegram)
	cl := chgk_api_client.New("https://db.chgk.info/")
	m := make(map[uint64]*chgk_api_client.Question)
	return &chgkServer{
		pb.UnimplementedChgkServiceServer{},
		tg,
		cl,
		repo,
		m,
	}
}

func (s chgkServer) SendMessage(ctx context.Context, in *pb.SendMessageReq) (*pb.Message, error) {
	mes, err := s.tg.SendMessage(in.ChatId, in.Text)
	if err != nil {
		return nil, err
	}
	return mes, nil
}

func (s chgkServer) GetRandomQuestion(ctx context.Context, in *pb.SendMessageReq) (*pb.Question, error) {
	q, err := s.chgk.GetRandomQuestion()
	if err != nil {
		return nil, err
	}

	question := &pb.Question{
		Id:       q.ID,
		Question: q.Question,
		Answer:   q.Answer,
		Authors:  q.Authors,
		Comments: q.Comments,
	}
	s.tg.SendMessage(in.ChatId, question.Question)
	return question, nil

}
