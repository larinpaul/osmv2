package main

import "net/http"

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("POST /api/customers/{customerID}/orders", h)
}
