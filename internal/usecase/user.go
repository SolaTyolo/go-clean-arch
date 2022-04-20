package usecase

import (
	"context"

	"github.com/SolaTyolo/go-clean-arch/internal/entity"
	"github.com/SolaTyolo/go-clean-arch/internal/usecase/repo"
)

type UserUseCase struct {
	repo repo.UserRepo
	// webApi UserWebApi  - link to third part server
}

func NewUserUseCase(r repo.UserRepo) *UserUseCase {
	return &UserUseCase{
		repo: r,
		// webApi: w,
	}
}

func (uc *UserUseCase) User(ctx context.Context, uid string) (entity.User, error) {
	u, _ := uc.repo.GetUser(ctx, uid)
	return *u, nil
}

func (uc *UserUseCase) SaveUser(ctx context.Context, t entity.User) error {
	err := uc.repo.SaveUser(ctx, &t)
	// u, err := uc.webApi.getOtherData(t)
	// biz logic .....
	return err
}
