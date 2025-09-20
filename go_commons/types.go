package main

import (
	"context"

	pb "github.com/sizokonpc/commons/api"
)

type OrdersService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, *pb.CreateOrderRequest) error
}

type OrdersStore interface {
	CreateOrder(context.Context) error
}
