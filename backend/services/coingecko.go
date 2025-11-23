package services

import (
	"crypto-marketplace/models"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"
)

const (
	coinGeckoBaseURL = "https://api.coingecko.com/api/v3"
	requestTimeout   = 10 * time.Second
)

type CoinGeckoService struct {
	client      *http.Client
	rateLimiter <-chan time.Time
}

// CoinGecko API response structures
type CoinGeckoMarketData struct {
	ID                string  `json:"id"`
	Symbol            string  `json:"symbol"`
	Name              string  `json:"name"`
	CurrentPrice      float64 `json:"current_price"`
	MarketCap         float64 `json:"market_cap"`
	TotalVolume       float64 `json:"total_volume"`
	PriceChange24h    float64 `json:"price_change_percentage_24h"`
	Image             string  `json:"image"`
	CirculatingSupply float64 `json:"circulating_supply"`
}

func NewCoinGeckoService() *CoinGeckoService {
	return &CoinGeckoService{
		client: &http.Client{
			Timeout: requestTimeout,
		},
		// Rate limiter: 30 calls per minute = 1 call every 2 seconds
		rateLimiter: time.Tick(2 * time.Second),
	}
}

// GetMarketData fetches real-time market data for specified cryptocurrencies
func (s *CoinGeckoService) GetMarketData(coins []string) ([]models.Product, error) {
	<-s.rateLimiter // Wait for rate limiter

	// Build the API URL
	coinIDs := strings.Join(coins, ",")
	url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&ids=%s&order=market_cap_desc&sparkline=false&price_change_percentage=24h",
		coinGeckoBaseURL, coinIDs)

	// Make the request
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch market data: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	// Parse the response
	var marketData []CoinGeckoMarketData
	if err := json.NewDecoder(resp.Body).Decode(&marketData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	// Convert to Product models
	products := make([]models.Product, 0, len(marketData))
	for _, data := range marketData {
		products = append(products, models.Product{
			ID:          data.ID,
			Name:        data.Name,
			Symbol:      strings.ToUpper(data.Symbol),
			Price:       data.CurrentPrice,
			Change24h:   data.PriceChange24h,
			Volume24h:   data.TotalVolume,
			MarketCap:   data.MarketCap,
			Description: fmt.Sprintf("%s (%s) - Real-time data from CoinGecko", data.Name, strings.ToUpper(data.Symbol)),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	return products, nil
}

// GetTopCoins fetches the top N cryptocurrencies by market cap
func (s *CoinGeckoService) GetTopCoins(limit int) ([]models.Product, error) {
	<-s.rateLimiter // Wait for rate limiter

	url := fmt.Sprintf("%s/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=%d&page=1&sparkline=false&price_change_percentage=24h",
		coinGeckoBaseURL, limit)

	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch top coins: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("API returned status %d: %s", resp.StatusCode, string(body))
	}

	var marketData []CoinGeckoMarketData
	if err := json.NewDecoder(resp.Body).Decode(&marketData); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	products := make([]models.Product, 0, len(marketData))
	for _, data := range marketData {
		products = append(products, models.Product{
			ID:          data.ID,
			Name:        data.Name,
			Symbol:      strings.ToUpper(data.Symbol),
			Price:       data.CurrentPrice,
			Change24h:   data.PriceChange24h,
			Volume24h:   data.TotalVolume,
			MarketCap:   data.MarketCap,
			Description: fmt.Sprintf("%s (%s) - Real-time data from CoinGecko", data.Name, strings.ToUpper(data.Symbol)),
			CreatedAt:   time.Now(),
			UpdatedAt:   time.Now(),
		})
	}

	log.Printf("Fetched %d coins from CoinGecko", len(products))
	return products, nil
}

// UpdatePrices fetches updated prices for existing products
func (s *CoinGeckoService) UpdatePrices(productIDs []string) (map[string]models.Product, error) {
	products, err := s.GetMarketData(productIDs)
	if err != nil {
		return nil, err
	}

	result := make(map[string]models.Product)
	for _, product := range products {
		result[product.ID] = product
	}

	return result, nil
}
