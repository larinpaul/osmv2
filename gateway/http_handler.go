package main

import (
	"errors"
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

	var items []*pb.ItemsWithQuality
	if err := common.ReadJSON(r, &items); err != nill {
		common.WriteError(w, http.StatusBadRequest, err.Error())
		return
	}

	err := validateItems(items)

	h.client.CreateOrder(r.Context(), &pb.CreateOrderRequest{
		customerID: customerID,
		Items:      items,
	})
}

func validateItems(items []*pb.ItemsWithQuality) error {
	if len(items) == 0 {
		return errors.New("items must have at least one items")
	}

}
