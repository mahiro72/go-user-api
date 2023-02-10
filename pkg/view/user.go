package view

import (
	"time"

	"github.com/mahiro72/go-user-api/pkg/domain/model"
)

type User struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUser(m *model.User) *User {
	return &User{
		Id:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}
