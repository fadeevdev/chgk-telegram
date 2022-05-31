package chgk

import (
	"context"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

func (s *chgkServer) CreateUser(ctx context.Context, req *pb.User) (*pb.ID, error) {

	var user = models.User{
		ID:        req.Id,
		Username:  req.Username,
		FirstName: req.FirstName,
		IsBot:     req.IsBot,
	}

	userID, err := s.repo.CreateUser(ctx, user)

	return &pb.ID{Id: userID}, err
}
