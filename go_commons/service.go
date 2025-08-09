package main

type service struct {
	store OrdersStore
}

func NewService(store OrdersStore) *service {
	return &service{store}
}
