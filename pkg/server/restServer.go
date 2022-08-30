package server

import (
	"context"
	"fmt"
	proto "github.com/Geniuskaa/job-task/pkg/gen/proto/v1"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"log"
	"net"
	"net/http"
)

const defaultPort = "8888"
const defaultHost = "0.0.0.0"

func NewRestServer(ctx context.Context, grpcSrv *gRPCServer) *http.Server {

	mux := runtime.NewServeMux()

	err := proto.RegisterRusProfileHandlerServer(ctx, mux, grpcSrv.companySrv)

	if err != nil {
		log.Panic(fmt.Errorf("NewRestServer failed: %w", err))
	}

	httpSrv := http.Server{
		Addr:    fmt.Sprintf("%s:%s", defaultHost, defaultPort),
		Handler: mux,
		BaseContext: func(listener net.Listener) context.Context {
			return ctx
		},
	}

	return &httpSrv
}
