package user

import (
	"context"

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
	user, err := u.repository.Get(ctx, input.Id)
	if err != nil {
		return nil, err
	}
	out := &GetOutput{
		User: user,
	}
	return out, nil
}

func (u *Usecase) Create(ctx context.Context, input *AddInput) (*AddOutput, error) {
	return nil, nil
}
