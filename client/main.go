package main

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/edmarfelipe/grpc-load-balancing/client/nameresolver"
	"github.com/edmarfelipe/grpc-load-balancing/shared"
	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func main() {
	cc, err := startClient()
	if err != nil {
		panic(err)
	}
	defer cc.Close()

	makeRPCs(cc, 100)
}

func startClient() (*grpc.ClientConn, error) {
	resolver.Register(nameresolver.NewBuilder([]string{"localhost:1101", "localhost:1102"}))

	cc, err := grpc.Dial(
		nameresolver.BuildURI(),
		grpc.WithDefaultServiceConfig(`{"loadBalancingConfig": [{"round_robin":{}}]}`),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
	if err != nil {
		return nil, fmt.Errorf("could not connect: %w", err)
	}
	return cc, nil
}

func makeRPCs(cc *grpc.ClientConn, n int) {
	usrService := shared.NewUserClient(cc)
	for i := 0; i < n; i++ {
		result, err := usrService.Hello(context.Background(), &shared.Request{Id: uuid.NewString()})
		if err != nil {
			slog.Error("Error while calling Hello RPC", "err", err)
		}
		slog.Info("Reply from the server: " + result.Message)
	}
}
