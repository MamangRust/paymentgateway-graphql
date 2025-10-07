package main

import (
	"log"

	"github.com/MamangRust/paymentgatewaygraphql/internal/app"
)

func main() {
	srv, err := app.NewServer()
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}

	if err := srv.Run(); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
