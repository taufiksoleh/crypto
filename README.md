# Crypto Marketplace

A real-time cryptocurrency marketplace with WebSocket support for live product monitoring.

## Architecture

This project consists of two separate applications:

- **Backend**: Golang REST API with WebSocket support
- **Frontend**: Nuxt.js 3 web application

## Features

- ✅ Real-time product updates via WebSocket
- ✅ CRUD operations for cryptocurrency products
- ✅ Responsive UI with modern design
- ✅ Live connection status indicator
- ✅ Auto-reconnection on WebSocket disconnect
- ✅ RESTful API with CORS support

## Project Structure

```
crypto-marketplace/
├── backend/          # Golang backend
│   ├── handlers/     # HTTP and WebSocket handlers
│   ├── models/       # Data models
│   ├── websocket/    # WebSocket hub and client
│   ├── main.go       # Entry point
│   └── go.mod        # Go dependencies
│
└── frontend/         # Nuxt.js frontend
    ├── assets/       # CSS and static assets
    ├── components/   # Vue components
    ├── composables/  # API and WebSocket composables
    ├── pages/        # Page components
    ├── app.vue       # Root component
    └── package.json  # Node dependencies
```

## Quick Start

### Backend

```bash
cd backend
go mod download
go run main.go
```

The backend will start on `http://localhost:8080`

- REST API: `http://localhost:8080/api`
- WebSocket: `ws://localhost:8080/ws`

### Frontend

```bash
cd frontend
npm install
npm run dev
```

The frontend will start on `http://localhost:3000`

## API Endpoints

### REST API

- `GET /api/products` - Get all products
- `GET /api/products/{id}` - Get a specific product
- `POST /api/products` - Create a new product
- `PUT /api/products/{id}` - Update a product
- `DELETE /api/products/{id}` - Delete a product

### WebSocket

Connect to `ws://localhost:8080/ws` to receive real-time updates.

**Update Message Format:**

```json
{
  "type": "price_update" | "new_product" | "delete_product",
  "product": {
    "id": "uuid",
    "name": "Bitcoin",
    "symbol": "BTC",
    "price": 45000.50,
    "change24h": 2.5,
    "volume24h": 28000000000,
    "marketCap": 880000000000,
    "description": "Description",
    "createdAt": "2024-01-01T00:00:00Z",
    "updatedAt": "2024-01-01T00:00:00Z"
  }
}
```

## Technology Stack

### Backend
- Go 1.21+
- Gorilla Mux (HTTP router)
- Gorilla WebSocket
- CORS middleware

### Frontend
- Nuxt.js 3
- Vue 3
- TypeScript
- Native WebSocket API

## Development

### Backend Development

The backend includes sample cryptocurrency data for testing. You can modify the `initSampleProducts()` function in `backend/handlers/product.go` to change the initial data.

### Frontend Development

The frontend is configured to connect to the backend at:
- API: `http://localhost:8080/api`
- WebSocket: `ws://localhost:8080/ws`

You can change these URLs in `frontend/nuxt.config.ts` or via environment variables:
- `API_BASE_URL`
- `WS_BASE_URL`

## License

MIT