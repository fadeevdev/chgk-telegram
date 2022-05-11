package chgk

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

type Repository interface {
	CreateUser(context.Context, models.User) (int, error)
}
