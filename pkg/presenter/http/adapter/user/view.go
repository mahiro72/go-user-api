package user

import (
	"time"

	"github.com/mahiro72/go-user-api/pkg/domain/model"
)

type UserView struct {
	Id        int       `json:"id"`
	Name      string    `json:"name"`
	CreatedAt time.Time `json:"created_at"`
}

func NewUserView(m *model.User) *UserView {
	return &UserView{
		Id:        m.ID,
		Name:      m.Name,
		CreatedAt: m.CreatedAt,
	}
}
