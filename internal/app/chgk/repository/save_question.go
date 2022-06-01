package repository

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk_api_client"
)

func (r *repository) SaveQuestion(ctx context.Context, question *chgk_api_client.Question) (ID uint64, err error) {
	const query = `
		insert into questions (
			id,
			question,
			answer,
			authors,
			comments
		) VALUES (
			$1, $2, $3, $4, $5
		) returning id
	`
	err = r.pool.QueryRow(ctx, query,
		question.ID,
		question.Question,
		question.Answer,
		question.Authors,
		question.Comments,
	).Scan(&ID)
	return
}
