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

func (r *repository) AddToTop(ctx context.Context, uID uint64, qID uint64) (err error) {
	const query = `
		update correct_answers (
			id,
			answered_questions
		) VALUES (
			$1, $2
		) ON CONFLICT (id) do
			update set answered_questions = array_prepend(answered_questions, $2);
	`
	r.pool.QueryRow(ctx, query,
		uID,
		qID,
	)

	return
}
