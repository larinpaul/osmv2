package main

import (
	"context"
	"log"
	"net"

	"google.go.golang.org/grpc"
)

var (
	grpcAddr = common.EnvString("GRCP_ADDR", "localhost:2000")
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Dial("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer f.Close()

	store := NewStore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())

	if err := grpcServer.Serve(); err != nil {
		log.Fatal(err.Error())
	}
}
