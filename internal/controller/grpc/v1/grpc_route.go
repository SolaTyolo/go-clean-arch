// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"github.com/SolaTyolo/go-clean-arch/internal/usecase"
	"github.com/SolaTyolo/go-clean-arch/pkg/grpcserver"
	"github.com/SolaTyolo/go-clean-arch/pkg/logger"
)

type GrpcRouteCase struct {
	UserUseCase usecase.UserUseCase
}

func NewGrpcRouter(s *grpcserver.Server, l logger.Logger, rc *GrpcRouteCase) {

	// grpc server
	us := newUserService(rc.UserUseCase, l)
	{
		RegisterUserServiceServer(s.Server(), us)
		l.Info("Registered user grpc server")
	}
	// append other service
}
