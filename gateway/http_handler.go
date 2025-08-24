package main

import (
	"net/http"
)

type handler struct {
	client pb.OrderServiceClient
}

func NewHandler(client pb.OrderServiceClient) *handler {
	return &handler{client}
}

func (h *handler) registerRoutes(nux *http.ServeMux) {
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h.HandleCreateOrder())
}

func (h *handler) HandleCreateOrder(w http.ResponseWriter, r *http.Request) {
	customerID := r.PathValue("customerID")

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		customerID: customerID,
		Items:      items,
	})
}
