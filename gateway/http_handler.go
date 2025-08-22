package main

import (
	"net/http"

	pb "github.com/sikozonpc/commons/api"
)

type handler struct {
	// gateway
}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) registerRoutes(nux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder())
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	pb.NewOrderServiceClient()
}
