package main

import (
	"crypto-marketplace/handlers"
	"crypto-marketplace/websocket"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Initialize WebSocket hub
	hub := websocket.NewHub()
	go hub.Run()

	// Initialize handlers
	productHandler := handlers.NewProductHandler(hub)
	wsHandler := handlers.NewWebSocketHandler(hub)

	// Setup router
	r := mux.NewRouter()

	// API routes
	api := r.PathPrefix("/api").Subrouter()
	api.HandleFunc("/products", productHandler.GetProducts).Methods("GET")
	api.HandleFunc("/products/{id}", productHandler.GetProduct).Methods("GET")
	api.HandleFunc("/products", productHandler.CreateProduct).Methods("POST")
	api.HandleFunc("/products/{id}", productHandler.UpdateProduct).Methods("PUT")
	api.HandleFunc("/products/{id}", productHandler.DeleteProduct).Methods("DELETE")

	// WebSocket route
	r.HandleFunc("/ws", wsHandler.HandleConnection)

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	// Start price simulator (simulates real-time price updates)
	go simulatePriceUpdates(productHandler, hub)

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Printf("WebSocket available at ws://localhost%s/ws", port)
	log.Printf("REST API available at http://localhost%s/api", port)

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// simulatePriceUpdates simulates real-time cryptocurrency price changes
func simulatePriceUpdates(handler *handlers.ProductHandler, hub *websocket.Hub) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		// This is a simple simulation - in production, you'd connect to real crypto APIs
		log.Println("Simulating price updates...")
	}
}
