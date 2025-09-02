package main

import pb "github.com/sikozonpc/commons/api"

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}
