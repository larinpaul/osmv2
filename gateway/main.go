package main

import (
	"log"
	"net/http"

	common "github.com/sikozonpc/commons"
)

const (
	httpAddr = common.EnvString("HTTP_ADDR", ":8080")
)

func main() {
	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log Fatal("Failed to start http server")
	}
}
