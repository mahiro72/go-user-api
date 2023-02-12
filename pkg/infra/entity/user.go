package entity

import (
	"time"

	"github.com/mahiro72/go-user-api/pkg/domain/model"
)

type User struct {
	ID        int       `gorm:"id"`
	Name      string    `gorm:"name"`
	CreatedAt time.Time `gorm:"created_at"`
}

func (u *User) ToModel() *model.User {
	return &model.User{
		ID:        u.ID,
		Name:      u.Name,
		CreatedAt: u.CreatedAt,
	}
}

func NewUserFromModel(m *model.User) *User {
	return &User{
		ID:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}
