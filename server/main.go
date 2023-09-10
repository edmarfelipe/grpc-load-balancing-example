package main

import (
	"context"
	"flag"
	"fmt"
	"log/slog"
	"net"

	"github.com/edmarfelipe/grpc-load-balancing/shared"
	"google.golang.org/grpc"
)

func main() {
	numbPtr := flag.Int("port", 1100, "grpc server port")
	flag.Parse()

	srv := grpc.NewServer()
	shared.RegisterUserServer(srv, &UserServer{serverPort: *numbPtr})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", *numbPtr))
	if err != nil {
		panic(fmt.Errorf("could not connect: %w", err))
	}

	slog.Info("Server is running on port " + fmt.Sprintf(":%d", *numbPtr))
	srv.Serve(listener)
}

type UserServer struct {
	shared.UnimplementedUserServer
	serverPort int
}

func (us *UserServer) Hello(ctx context.Context, req *shared.Request) (*shared.Reply, error) {
	slog.Info("Recive Hello with ID: " + req.Id)
	return &shared.Reply{
		Message: fmt.Sprintf("Hello port: %d", us.serverPort),
	}, nil
}
