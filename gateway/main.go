package main

import (
	"log"
	"net/http"

	_ "github.com/joho/gotoenv/autoload"
	common "github.com/sikozonpc/commons"
)

const (
	httpAddr = common.EnvString("HTTP_ADDR", ":3000") 
	orderServiceAddr = "localhost:3000"
)

func main() {
	conn, err := grpc.Dial(orderServiceAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to dial server: %v", err)
	}
	defer conn.Close()

	log.Println("Dialing orders service at ", ordersServiceAddr)

	mux := http.NewServeMux()
	handler := NewHandler()
	handler.registerRoutes(mux)

	log.Printf("Starting HTTP server at %s", httpAddr)

	if err := http.ListenAndServe(httpAddr, mux); err != nil {
		log Fatal("Failed to start http server")
	}
}
