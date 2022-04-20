// Package v1 implements routing paths. Each services in own file.
package v1

import (
	"context"
	"fmt"

	"github.com/SolaTyolo/go-clean-arch/pkg/grpcserver"
	"github.com/SolaTyolo/go-clean-arch/pkg/logger"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type GatewayRouteConfig struct {
	Gws *grpcserver.GatewayServer
	Gs  *grpcserver.Server
}

func NewGatewayRouter(c *GatewayRouteConfig, l logger.Logger) {

	// gateway call grpc server options
	opts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithDefaultCallOptions(
			grpc.MaxCallRecvMsgSize(10*1024*1024),
			grpc.MaxCallSendMsgSize(10*1024*1024),
		),
	}

	// 当前所有grpc-server使用同一个addr
	endpoint := fmt.Sprintf("localhost%s", c.Gs.Addr)

	// gateway server 定义追加在这边
	{
		err := RegisterUserServiceHandlerFromEndpoint(context.Background(), c.Gws.Mux, endpoint, opts)
		if err != nil {
			l.Fatal("failed to register user gateway handler %v", err)
		}
		l.Info("Registered user http server")
	}
	// append other service handle
}
