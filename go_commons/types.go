package main

import (
	"context"
)

type OrdersService interface {
	CreateOrder(context.Context) error
	ValidateOrder(context.Context, pb.) error
}

type OrdersStore interface {
	CreateOrder(context.Context) error
}
