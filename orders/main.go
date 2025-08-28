package main

import (
	"context"
	"log"
	"net"

	"google.go.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()

	l, err := net.Dial("tcp", grpcAddr)

	store := NewStore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())

	if err := grpcServer.Serve(); err != nil {
		log.Fatal(err.Error())
	}
}
