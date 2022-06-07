package chgk

import (
	"context"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
)

func (s *chgkServer) GetTopPlayers(ctx context.Context, u *pb.Count) (*pb.TopUsers, error) {
	return nil, nil
}

func (s *chgkServer) GetTopPosition(ctx context.Context, u *pb.User) (top *pb.TopUser, err error) {
	position, err := s.repo.GetTopPosition(ctx, u.Id)
	if err != nil {
		return
	}
	top.Position = position.Position
	top.FirstName = position.Username
	top.Questions = position.Questions

	return
}
