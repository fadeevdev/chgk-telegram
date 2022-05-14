package database

import (
	"context"
	"fmt"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk/config"

	"github.com/jackc/pgx/v4/pgxpool"
)

func ConnectionString(cfg *config.Postgres) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.Host, cfg.Port, cfg.DbUser, cfg.Password, cfg.DbName)
}

func NewPool(ctx context.Context, cfg *config.Postgres) (*pgxpool.Pool, error) {
	return pgxpool.Connect(ctx, ConnectionString(cfg))
}
