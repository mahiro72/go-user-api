package repository

import (
	"context"

	"github.com/mahiro72/go-user-api/pkg/domain/model"
)

type IUser interface {
	Get(ctx context.Context, id int) (*model.User, error)
	Create(ctx context.Context, m *model.User) (*model.User, error)
}
