package chgk

import (
	"context"
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
)

func (s *chgkServer) WebHook(ctx context.Context, message *pb.Message) (*pb.Empty, error) {
	switch message.Text {
	case "/start":
		fmt.Println("test")
		fmt.Println(s.RegisterUser(ctx, message.From))
	}
	return &pb.Empty{}, nil
}
