package main

import (
	"context"
	"log"
	"net"
)

var (
	grpcAddr = common.EnvString("GRCP_ADDR", "localhost:2000")
)

func main() {

	l, err := net.Listen("tcp", grpcAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	defer f.Close()

	store := NewStore()
	svc := NewService(store)
	NewGRPCHandler(grpcServer)

	svc.CreateOrder(context.Background())

	if err := grpcServer.Serve(l); err != nil {
		log.Fatal(err.Error())
	}
}
