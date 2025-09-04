package main

import (
	"context"
	"log"

	pb "github.com/sikozonpc/commons/api"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer) *grpcHandler {
	return &grpcHandler{grpcServer}
}

func (h *grpcHandler) CreateOrder(context.Context, *pb.CreteOrderRequest) (*pb.Order, error) {
	log.Println("New order received!")
	o := &pb.Order{
		Id: "42",
	}
	return o, nil
}
