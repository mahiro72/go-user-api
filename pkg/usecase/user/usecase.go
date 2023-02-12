package user

import (
	"context"
	"fmt"

	"github.com/mahiro72/go-user-api/pkg/domain/model"
	"github.com/mahiro72/go-user-api/pkg/domain/repository"
)

type Usecase struct {
	repository repository.IUser
}

func NewUsecase(repository repository.IUser) *Usecase {
	return &Usecase{
		repository: repository,
	}
}

func (u *Usecase) Get(ctx context.Context, input *GetInput) (*GetOutput, error) {
	user, err := u.repository.Get(ctx, input.ID)
	if err != nil {
		return nil, err
	}

	out := &GetOutput{
		User: user,
	}
	return out, nil
}

func (u *Usecase) Create(ctx context.Context, input *CreateInput) (*CreateOutput, error) {
	if input.Name == "" {
		return nil, fmt.Errorf("error: user name is empty")
	}

	m := &model.User{
		Name: input.Name,
	}
	user, err := u.repository.Create(ctx, m)
	if err != nil {
		return nil, err
	}

	out := &CreateOutput{
		User: user,
	}
	return out, nil
}
