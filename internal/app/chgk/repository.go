package chgk

import (
	"context"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/chgk_api_client"
	"gitlab.ozon.dev/fadeevdev/homework-2/internal/app/models"
)

type Repository interface {
	RegisterUser(context.Context, models.User) (uint64, error)
	GetUser(context.Context, uint64) (models.User, error)
	SaveQuestion(ctx context.Context, question *chgk_api_client.Question) (uint64, error)
}
