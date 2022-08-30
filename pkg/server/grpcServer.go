package server

import (
	"context"
	proto "github.com/Geniuskaa/job-task/pkg/gen/proto/v1"
	"github.com/Geniuskaa/job-task/pkg/gen/proto/v1/company"
	"google.golang.org/grpc"
)

type gRPCServer struct {
	Srv        *grpc.Server
	companySrv company.Service
}

func NewGRPCServer(ctx context.Context) *gRPCServer {
	grpcSrv := grpc.NewServer()

	return &gRPCServer{Srv: grpcSrv,
		companySrv: *company.NewService(),
	}
}

func (g *gRPCServer) Init() {
	proto.RegisterRusProfileServer(g.Srv, g.companySrv)
}

func (g *gRPCServer) GetSrv() *grpc.Server {
	return g.Srv
}
