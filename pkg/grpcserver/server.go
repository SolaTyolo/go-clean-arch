package grpcserver

import (
	"net"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	grpc_ctxtags "github.com/grpc-ecosystem/go-grpc-middleware/tags"
	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-prometheus"
	"go.elastic.co/apm/module/apmgrpc"

	"github.com/SolaTyolo/go-clean-arch/pkg/logger"

	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
)

const (
	_defaultMaxRecvMsgSize = 1024 * 1024 * 1024
	_defaultMaxSendMsgSize = 1024 * 1024 * 1024
)

// Server -
type Server struct {
	ServiceName string
	// grpc server
	gs         *grpc.Server
	notify     chan error
	Addr       string
	UseGateway bool
	log        *logger.Logger
}

type Config struct {
	AppName string
	Addr    string
}

func New(c *Config, log *logger.Logger) *Server {

	serve := &Server{
		ServiceName: c.AppName,
		notify:      make(chan error, 1),
		Addr:        c.Addr,
		log:         log,
	}

	// TODO - 配置化
	ms := 1024 * 1024 * 1024
	s := grpc.NewServer(
		grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
			apmgrpc.NewUnaryServerInterceptor(),
			grpc_prometheus.UnaryServerInterceptor,
			grpc_ctxtags.UnaryServerInterceptor(),
			grpc_recovery.UnaryServerInterceptor(),
		)),
		grpc.MaxRecvMsgSize(ms),
		grpc.MaxSendMsgSize(ms),
	)

	log.Info("started grpc server at addr %s", c.Addr)
	grpc_health_v1.RegisterHealthServer(s, health.NewServer())

	reflection.Register(s)
	grpc_prometheus.Register(s)
	grpc_prometheus.EnableHandlingTimeHistogram()

	serve.gs = s
	serve.start()
	return serve
}

func (s *Server) start() {
	lis, err := net.Listen("tcp", s.Addr)
	if err != nil {
		s.log.Fatal("failed to listen: %v", err)
	}
	go func() {
		s.notify <- s.gs.Serve(lis)
		close(s.notify)
	}()
}
func (s *Server) Server() *grpc.Server {
	return s.gs
}

func (s *Server) Notify() <-chan error {
	return s.notify
}

func (s *Server) Shutdown() error {
	s.gs.GracefulStop()
	s.log.Info("GrpcServer:%s ShutDown", s.ServiceName)
	return nil
}
