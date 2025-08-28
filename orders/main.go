package main

import (
	"context"

	"google.go.golang.org/grpc"
)

func main() {

	grpcServer := grpc.NewServer()

	store := NewStore()
	svc := NewService(store)

	svc.CreateOrder(context.Background())
}
