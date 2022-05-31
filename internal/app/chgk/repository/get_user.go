package repository

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

func (r *repository) GetUser(ctx context.Context, id uint64) (user models.User, err error) {

	const query = `
		select * from users
		where id = $1;
	`

	err = r.pool.QueryRow(ctx, query,
		id,
	).Scan(&user.ID, &user.Username, &user.FirstName, &user.IsBot, &user.CreatedAt)

	return
}
