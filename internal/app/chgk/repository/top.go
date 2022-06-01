package repository

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

func (r *repository) GetTopPlayers(ctx context.Context, n uint64) ([]models.User, error) {
	//TBD
	return nil, nil
}

func (r *repository) GetTopPosition(ctx context.Context, u models.User) ([]models.TopUser, error) {
	//TBD
	const query = `
		select * from questions
	`
	return nil, nil
}
