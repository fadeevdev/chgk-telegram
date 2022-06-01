package chgk

import (
	"context"
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
)

func (s *chgkServer) WebHook(ctx context.Context, message *pb.Message) (*pb.Empty, error) {
	fmt.Printf("received a message: %s\n", message)
	switch message.Text {
	case "/start":
		_, err := s.RegisterUser(ctx, message.From)
		if err != nil {
			return &pb.Empty{}, err
		}
	}
	return &pb.Empty{}, nil
}
