package v1

import (
	"context"

	"github.com/SolaTyolo/go-clean-arch/internal/entity"
	"github.com/SolaTyolo/go-clean-arch/internal/usecase"
	"github.com/SolaTyolo/go-clean-arch/pkg/logger"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

type UserService struct {
	usercase usecase.UserUseCase
	l        logger.Logger
}

func newUserService(uc usecase.UserUseCase, l logger.Logger) *UserService {
	return &UserService{uc, l}
}

// GetUser -
func (u *UserService) GetUser(ctx context.Context, req *GetUserRequest) (*User, error) {
	user, err := u.usercase.User(ctx, req.GetUserId())
	if err != nil {
		u.l.Error("grpc - v1 - getuser", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	return &User{Name: user.Name, UserType: user.Age, Id: user.ID, CreateDate: ""}, nil
}

// UpsertUser -
func (u *UserService) UpsertUser(ctx context.Context, req *UpsertUserInput) (*emptypb.Empty, error) {
	err := u.usercase.SaveUser(ctx, entity.User{Name: req.Name, Age: req.Age})
	if err != nil {
		u.l.Error("grpc - v1 - upsertuser", err)
		return nil, status.Errorf(codes.Internal, err.Error())
	}
	return &emptypb.Empty{}, nil
}
