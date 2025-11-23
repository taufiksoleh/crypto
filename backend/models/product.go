package models

import "time"

type Product struct {
	ID          string    `json:"id"`
	Name        string    `json:"name"`
	Symbol      string    `json:"symbol"`
	Price       float64   `json:"price"`
	Change24h   float64   `json:"change24h"`
	Volume24h   float64   `json:"volume24h"`
	MarketCap   float64   `json:"marketCap"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

type ProductUpdate struct {
	Type    string  `json:"type"` // "price_update", "new_product", "delete_product"
	Product Product `json:"product"`
}
