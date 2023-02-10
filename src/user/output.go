package user

import "github.com/mahiro72/go-user-api/pkg/domain/model"

type GetOutput struct {
	User *model.User
}

type AddOutput struct {
	User *model.User
}
