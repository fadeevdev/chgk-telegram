package chgk

import (
	"context"
	"fmt"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
)

func (s *chgkServer) WebHook(ctx context.Context, update *pb.Update) (*pb.Empty, error) {
	fmt.Printf("received a message: %s\n", update)
	switch update.Message.Text {
	case "/start":
		_, err := s.RegisterUser(ctx, update.Message.From)
		if err != nil {
			return &pb.Empty{}, err
		}
	}
	return &pb.Empty{}, nil
}
