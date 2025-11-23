package handlers

import (
	"crypto-marketplace/models"
	"crypto-marketplace/websocket"
	"encoding/json"
	"net/http"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

type ProductHandler struct {
	products map[string]models.Product
	mu       sync.RWMutex
	hub      *websocket.Hub
}

func NewProductHandler(hub *websocket.Hub) *ProductHandler {
	handler := &ProductHandler{
		products: make(map[string]models.Product),
		hub:      hub,
	}

	// Initialize with some sample products
	handler.initSampleProducts()

	return handler
}

func (h *ProductHandler) initSampleProducts() {
	sampleProducts := []models.Product{
		{
			ID:          uuid.New().String(),
			Name:        "Bitcoin",
			Symbol:      "BTC",
			Price:       45000.50,
			Change24h:   2.5,
			Volume24h:   28000000000,
			MarketCap:   880000000000,
			Description: "The first and most valuable cryptocurrency",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Name:        "Ethereum",
			Symbol:      "ETH",
			Price:       3200.75,
			Change24h:   -1.2,
			Volume24h:   15000000000,
			MarketCap:   385000000000,
			Description: "Leading smart contract platform",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
		{
			ID:          uuid.New().String(),
			Name:        "Cardano",
			Symbol:      "ADA",
			Price:       0.58,
			Change24h:   3.8,
			Volume24h:   550000000,
			MarketCap:   20500000000,
			Description: "Proof-of-stake blockchain platform",
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		},
	}

	for _, product := range sampleProducts {
		h.products[product.ID] = product
	}
}

func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	h.mu.RLock()
	defer h.mu.RUnlock()

	products := make([]models.Product, 0, len(h.products))
	for _, product := range h.products {
		products = append(products, product)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.mu.RLock()
	product, exists := h.products[id]
	h.mu.RUnlock()

	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product models.Product
	if err := json.NewDecoder(r.Body).Decode(&product); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	product.ID = uuid.New().String()
	product.CreatedAt = time.Now()
	product.UpdatedAt = time.Now()

	h.mu.Lock()
	h.products[product.ID] = product
	h.mu.Unlock()

	// Broadcast new product via WebSocket
	h.hub.BroadcastUpdate(models.ProductUpdate{
		Type:    "new_product",
		Product: product,
	})

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(product)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.mu.Lock()
	defer h.mu.Unlock()

	existingProduct, exists := h.products[id]
	if !exists {
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	var updates models.Product
	if err := json.NewDecoder(r.Body).Decode(&updates); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Update fields
	existingProduct.Name = updates.Name
	existingProduct.Symbol = updates.Symbol
	existingProduct.Price = updates.Price
	existingProduct.Change24h = updates.Change24h
	existingProduct.Volume24h = updates.Volume24h
	existingProduct.MarketCap = updates.MarketCap
	existingProduct.Description = updates.Description
	existingProduct.UpdatedAt = time.Now()

	h.products[id] = existingProduct

	// Broadcast price update via WebSocket
	h.hub.BroadcastUpdate(models.ProductUpdate{
		Type:    "price_update",
		Product: existingProduct,
	})

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(existingProduct)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	h.mu.Lock()
	product, exists := h.products[id]
	if !exists {
		h.mu.Unlock()
		http.Error(w, "Product not found", http.StatusNotFound)
		return
	}

	delete(h.products, id)
	h.mu.Unlock()

	// Broadcast deletion via WebSocket
	h.hub.BroadcastUpdate(models.ProductUpdate{
		Type:    "delete_product",
		Product: product,
	})

	w.WriteHeader(http.StatusNoContent)
}
