package main

import (
	"context"
	"log"

	pb "github.com/sikozonpc/commons/api"
	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer

	service OrdersService
}

func NewGRPCHandler(grpcServer *grpc.Server) *grpcHandler {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer{grpcServer, handler}
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreteOrderRequest) (*pb.Order, error) {
	log.Printf("New order received! Order %v", p)
	o := &pb.Order{
		Id: "42",
	}
	return o, nil
}
