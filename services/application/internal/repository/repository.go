package repository

import (
	"context"

	"github.com/lucas-hill/credit-decision/services/application/internal/model"
)

type ApplicationRepository interface {
	Create(ctx context.Context, app *model.Application) error
	GetByID(ctx context.Context, id string) (*model.Application, error)
}
