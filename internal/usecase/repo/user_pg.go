package repo

import (
	"context"

	"github.com/SolaTyolo/go-clean-arch/internal/entity"
	"github.com/SolaTyolo/go-clean-arch/pkg/postgres"
)

type UserRepo struct {
	*postgres.Postgres
}

func NewUserRepo(pg *postgres.Postgres) *UserRepo {
	return &UserRepo{pg}
}

func (r *UserRepo) GetUser(ctx context.Context, id string) (*entity.User, error) {
	var u entity.User

	tx := r.DB.WithContext(ctx)
	tx = tx.Where("id = ?", id)
	if err := tx.Find(&u).Error; err != nil {
		return nil, err
	}
	return &u, nil

}

func (r *UserRepo) SaveUser(ctx context.Context, e *entity.User) error {
	return r.DB.WithContext(ctx).Create(&e).Error
}
