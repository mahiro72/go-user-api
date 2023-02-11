package registry

import (
	"gorm.io/gorm"

	"github.com/mahiro72/go-user-api/pkg/db/mysql"
	"github.com/mahiro72/go-user-api/pkg/infra/dao"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository() (*Repository, error) {
	db, err := mysql.Init()
	if err != nil {
		return nil, err
	}

	return &Repository{db: db}, nil
}

func (r *Repository) NewUser() *dao.User {
	return dao.NewUser(r.db)
}
