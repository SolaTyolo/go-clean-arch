// Package httpserver implements HTTP server.
package grpcserver

import (
	"net/http"

	"github.com/SolaTyolo/go-clean-arch/pkg/httpserver.go"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
)

// Server -.
type GatewayServer struct {
	Httpserver *httpserver.Server
	Mux        *runtime.ServeMux
}

// New -.
func NewGatewayServer(gatewayName string, opts ...httpserver.Option) *GatewayServer {

	mux := runtime.NewServeMux()
	httpMux := http.NewServeMux()
	httpMux.Handle("/", mux)

	hs := httpserver.New(httpMux, opts...)

	s := &GatewayServer{
		Httpserver: hs,
		Mux:        mux,
	}
	return s
}
