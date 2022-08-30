package main

import (
	"context"
	"flag"
	"fmt"
	proto "github.com/Geniuskaa/job-task/pkg/gen/proto/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"os"
)

const defaultPort = "9997"
const defaultHost = "0.0.0.0"

func main() {

	inn := flag.String("inn", "7714402621", "INN need to send valid request")
	flag.Parse()

	port, ok := os.LookupEnv("APP_PORT")
	if !ok {
		port = defaultPort
	}

	host, ok := os.LookupEnv("APP_HOST")
	if !ok {
		host = defaultHost
	}

	if err := execute(net.JoinHostPort(host, port), *inn); err != nil {
		log.Print(err)
		os.Exit(1)
	}
}

func execute(addr string, inn string) (err error) {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return err
	}
	defer func() {
		if cerr := conn.Close(); cerr != nil {
			if err == nil {
				err = cerr
				return
			}
			log.Print(err)
		}
	}()

	err = getCompanyInfo(conn, inn)
	if err != nil {
		return err
	}

	return nil
}

// TODO: через консоль вставить аргумент
func getCompanyInfo(conn *grpc.ClientConn, inn string) error {
	client := proto.NewRusProfileClient(conn)
	ctx := context.Background()

	msg, err := client.GetCompanyInfo(ctx, &proto.CompanyINN{Inn: inn})
	if err != nil {
		log.Print(fmt.Errorf("getCompanyInfo failed: %w", err))
		return err
	}

	fmt.Println("Server response: \n", msg)
	return nil
}
