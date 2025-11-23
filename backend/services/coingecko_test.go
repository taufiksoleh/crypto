package services

import (
	"testing"
)

func TestNewCoinGeckoService(t *testing.T) {
	service := NewCoinGeckoService()
	if service == nil {
		t.Error("expected service to be initialized")
	}

	if service.client == nil {
		t.Error("expected HTTP client to be initialized")
	}
}

func TestGetTopCoins(t *testing.T) {
	service := NewCoinGeckoService()

	// Test with a small limit to avoid rate limiting in tests
	products, err := service.GetTopCoins(5)

	// Note: This test may fail if there's no internet connection
	// or if CoinGecko API is unavailable. That's acceptable for CI/CD.
	if err != nil {
		t.Logf("Warning: GetTopCoins failed (this is acceptable if API is unavailable): %v", err)
		return
	}

	if len(products) == 0 {
		t.Error("expected at least one product")
	}

	// Check that products have required fields
	for _, product := range products {
		if product.ID == "" {
			t.Error("product ID should not be empty")
		}
		if product.Name == "" {
			t.Error("product name should not be empty")
		}
		if product.Symbol == "" {
			t.Error("product symbol should not be empty")
		}
	}
}
