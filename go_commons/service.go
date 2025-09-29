package main

import (
	"context"

	common "command-line-argumentsC:\\programming\\golang\\osmv2\\common\\errors.go"
	pb "github.com/sizokonpc/commons/api"
)

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}

func (s *service) CreateOrder(context.Context) {
	return nil
}

func (s *service) ValidateOrder(ctx context.Context, p *pb.CreateOrderRequest) error {
	if len(p.Items) == 0 {
		return common.ErrNoItems
	}

	mergedItems := mergeItemsQuantities(p.Items)
}

func mergeItemsQuantities(items []*pb.ItemsWithQuantity) []*pb.ItemsWithQuantity {

}
