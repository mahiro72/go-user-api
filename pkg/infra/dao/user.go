package dao

import (
	"context"

	"gorm.io/gorm"

	"github.com/mahiro72/go-user-api/pkg/domain/model"
	"github.com/mahiro72/go-user-api/pkg/infra/entity"
)

type User struct {
	db *gorm.DB
}

func NewUser(db *gorm.DB) *User {
	return &User{
		db: db,
	}
}

func (d *User) Get(ctx context.Context, id int) (*model.User, error) {
	var e entity.User
	if err := d.db.Where("id = ?", id).First(&e).Error; err != nil {
		return nil, err
	}

	return e.ToModel(), nil
}

func (d *User) Create(ctx context.Context, m *model.User) (*model.User, error) {
	e := entity.NewUserFromModel(m)
	if err := d.db.Create(e).Error; err != nil {
		return nil, err
	}

	return e.ToModel(), nil
}
