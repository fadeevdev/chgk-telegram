package chgk

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

type Repository interface {
	RegisterUser(context.Context, models.User) (uint64, error)
	GetUser(context.Context, uint64) (models.User, error)
}
