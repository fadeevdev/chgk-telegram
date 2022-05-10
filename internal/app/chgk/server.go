package chgk

import (
	"context"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/telegram"
)

type chgkServer struct {
	pb.UnimplementedChgkServiceServer
	tg *telegram.Client
}

func New(config *config.Config) *chgkServer {
	tg := telegram.New(config.ApiKeys.Telegram)
	return &chgkServer{
		pb.UnimplementedChgkServiceServer{},
		tg,
	}
}

func (s chgkServer) Register(ctx context.Context, in *pb.User) (*pb.Error, error) {
	return &pb.Error{
		Exists:  false,
		Message: "",
	}, nil
}

func (s chgkServer) SendMessage(ctx context.Context, in *pb.SendMessageReq) (*pb.Empty, error) {
	err := s.tg.SendMessage(in.ChatId, in.Message)
	if err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}
