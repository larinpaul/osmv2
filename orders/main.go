package main

import (
	"context"
	"log"

	"google.go.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()

	store := NewStore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())

	if err := grpcServer.Serve(); err != nil {
		log.Fatal(err.Error())
	}
}
