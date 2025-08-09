package main

import "context"

type store struct {
	// add here our mongoDB instance
}

func NewStore() *store { // constructor for the store
	return &store{}
}

func (s *store) Create(context.Context) {
	return nil
}
