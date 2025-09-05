package main

import (
	"context"
	"log"

	pb "github.com/sikozonpc/commons/api"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) *grpcHandler {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer{grpcServer, handler}
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreteOrderRequest) (*pb.Order, error) {
	log.Println("New order received!")
	o := &pb.Order{
		Id: "42",
	}
	return o, nil
}
