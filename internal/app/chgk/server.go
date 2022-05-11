package chgk

import (
	"context"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/telegram"
)

type chgkServer struct {
	pb.UnimplementedChgkServiceServer
	tg   *telegram.Client
	repo Repository
}

func New(config *config.Config, repo Repository) *chgkServer {
	tg := telegram.New(config.ApiKeys.Telegram)
	return &chgkServer{
		pb.UnimplementedChgkServiceServer{},
		tg,
		repo,
	}
}

func (s chgkServer) Register(ctx context.Context, in *pb.User) (*pb.Empty, error) {
	return &pb.Empty{}, nil
}

func (s chgkServer) SendMessage(ctx context.Context, in *pb.SendMessageReq) (*pb.Message, error) {
	mes, err := s.tg.SendMessage(in.ChatId, in.Text)
	if err != nil {
		return nil, err
	}
	return mes, nil
}
