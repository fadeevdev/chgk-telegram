package chgk

import (
	"context"
	"errors"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

const ErrUserAlreadyExists = "user already exists"

func (s *chgkServer) Start(ctx context.Context, req *pb.User) (*pb.ID, error) {

	u, _ := s.GetUser(ctx, req.Id)

	if u != nil && u.Id == req.Id {
		return &pb.ID{Id: req.Id}, errors.New(ErrUserAlreadyExists)
	}

	var user = models.User{
		ID:        req.Id,
		Username:  req.Username,
		FirstName: req.FirstName,
		IsBot:     req.IsBot,
	}

	userID, err := s.repo.Start(ctx, user)

	return &pb.ID{Id: userID}, err
}
