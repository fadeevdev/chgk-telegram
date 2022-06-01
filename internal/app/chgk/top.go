package chgk

import (
	"context"
	pb "gitlab.ozon.dev/fadeevdev/homework-2/api"
)

func (s *chgkServer) GetTopPlayers(ctx context.Context, u *pb.Count) (*pb.TopUsers, error) {
	return nil, nil
}

func (s *chgkServer) GetTopPosition(ctx context.Context, u *pb.User) (*pb.TopUser, error) {
	return nil, nil
}
