package main

import (
	"crypto-marketplace/handlers"
	"crypto-marketplace/services"
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

	// Initialize CoinGecko service
	coinGeckoService := services.NewCoinGeckoService()

	// Initialize handlers
	productHandler := handlers.NewProductHandler(hub, coinGeckoService)
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

	// Health check endpoint
	r.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(`{"status":"healthy","service":"crypto-marketplace"}`))
	}).Methods("GET")

	// CORS configuration
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://localhost:3001"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
	})

	handler := c.Handler(r)

	// Start real-time price updater (fetches from CoinGecko every 30 seconds)
	go updateRealTimePrices(productHandler)

	// Start server
	port := ":8080"
	log.Printf("Server starting on port %s", port)
	log.Printf("WebSocket available at ws://localhost%s/ws", port)
	log.Printf("REST API available at http://localhost%s/api", port)

	if err := http.ListenAndServe(port, handler); err != nil {
		log.Fatal("Server failed to start:", err)
	}
}

// updateRealTimePrices fetches real-time price updates from CoinGecko
func updateRealTimePrices(handler *handlers.ProductHandler) {
	// Wait 30 seconds before first update to respect rate limits
	ticker := time.NewTicker(30 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		log.Println("Fetching real-time price updates from CoinGecko...")

		if err := handler.UpdatePricesFromCoinGecko(); err != nil {
			log.Printf("Error updating prices: %v", err)
		} else {
			log.Println("Successfully updated cryptocurrency prices")
		}
	}
}
