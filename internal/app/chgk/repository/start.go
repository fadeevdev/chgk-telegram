package repository

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

func (r *repository) Start(ctx context.Context, user models.User) (ID uint64, err error) {

	const query = `
		insert into users (
			id,
			username,
			firstname,
			is_bot,
			created_at
		) VALUES (
			$1, $2, $3, $4, now()
		) returning id
	`

	err = r.pool.QueryRow(ctx, query,
		user.ID,
		user.Username,
		user.FirstName,
		user.IsBot,
	).Scan(&ID)

	return
}
