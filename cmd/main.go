package main

import (
	"context"
	"fmt"
	"github.com/Geniuskaa/job-task/pkg/server"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

const (
	defaultPort = "9997"
	defaultHost = "0.0.0.0"
)

func main() {
	mainCtx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGHUP, syscall.SIGINT,
		syscall.SIGQUIT, syscall.SIGTERM)
	defer stop()

	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	grpcSrv, gRPClistener, httpSrv := applicationStart(mainCtx, net.JoinHostPort(host, port))

	g, gCtx := errgroup.WithContext(mainCtx)
	g.Go(func() error {
		fmt.Println("HTTP server started!")
		return httpSrv.ListenAndServe()
	})
	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("HTTP server is shut down.")
		return httpSrv.Shutdown(context.Background())
	})
	g.Go(func() error {
		fmt.Println("gRPC server started!")
		return grpcSrv.Serve(gRPClistener)
	})
	g.Go(func() error {
		<-gCtx.Done()
		fmt.Println("gRPC server is shut down.")
		grpcSrv.GracefulStop()
		return nil
	})

	if err := g.Wait(); err != nil {
		fmt.Printf("exit reason: %s \n", err)
	}
	fmt.Println("Servers were gracefully shut down.")
}

func applicationStart(ctx context.Context, addr string) (*grpc.Server, net.Listener, *http.Server) {
	gRPClistener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}

	grpcServer := server.NewGRPCServer(ctx)
	grpcServer.Init()

	restServer := server.NewRestServer(ctx, grpcServer)

	return grpcServer.GetSrv(), gRPClistener, restServer
}
