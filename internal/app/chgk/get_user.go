package chgk

import (
	"context"
	"errors"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
)

const ErrUserNotFound = "user not found"

func (s *chgkServer) GetUser(ctx context.Context, id uint64) (*pb.User, error) {

	u, err := s.repo.GetUser(ctx, id)

	if err != nil {
		return nil, err
	}

	var user = pb.User{
		Id:        u.ID,
		Username:  u.Username,
		FirstName: u.FirstName,
		IsBot:     u.IsBot,
	}

	if u.ID != id {
		return &user, errors.New(ErrUserNotFound)
	}

	return &user, err
}
