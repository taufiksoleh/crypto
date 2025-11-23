# Crypto Marketplace

A real-time cryptocurrency marketplace with WebSocket support for live product monitoring.

## Architecture

This project consists of two separate applications:

- **Backend**: Golang REST API with WebSocket support
- **Frontend**: Nuxt.js 3 web application

## Features

- ✅ **Real-time Cryptocurrency Data** - Powered by CoinGecko API
- ✅ **Live Price Updates** - Auto-updates every 30 seconds via WebSocket
- ✅ **Top 20 Cryptocurrencies** - Displays leading cryptocurrencies by market cap
- ✅ **CRUD Operations** - Create, read, update, and delete cryptocurrency products
- ✅ **Responsive UI** - Modern design with live connection status
- ✅ **Auto-reconnection** - WebSocket automatically reconnects on disconnect
- ✅ **RESTful API** - Clean API with CORS support
- ✅ **CI/CD Pipeline** - GitHub Actions for automated testing and deployment
- ✅ **Docker Support** - Production and development Docker configurations

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
- CoinGecko API (Real-time crypto data)
- CORS middleware

### Frontend
- Nuxt.js 3
- Vue 3
- TypeScript
- Native WebSocket API

## Real-time Data Integration

This marketplace uses the **CoinGecko API** to fetch real-time cryptocurrency data:

- **Data Source**: [CoinGecko API](https://www.coingecko.com/en/api)
- **Update Frequency**: Every 30 seconds
- **Data Points**: Price, 24h change, volume, market cap
- **Rate Limit**: 30 calls per minute (free tier)
- **No Authentication Required**: Free public endpoints

The backend automatically:
1. Loads the top 20 cryptocurrencies on startup
2. Updates prices every 30 seconds from CoinGecko
3. Broadcasts updates to all connected WebSocket clients
4. Falls back to sample data if the API is unavailable

## Using Docker

### Development Environment

```bash
# Start with hot-reloading
make dev

# Or manually
docker-compose -f docker-compose.dev.yml up
```

### Production Environment

```bash
# Build and start
make prod

# Or manually
docker-compose up --build -d

# View logs
docker-compose logs -f

# Stop services
docker-compose down
```

### Available Make Commands

```bash
make help          # Show all available commands
make dev           # Start development environment
make prod          # Start production environment
make test          # Run all tests
make lint          # Run linters
make clean         # Clean build artifacts
make health        # Check service health
```

## CI/CD Pipeline

The project includes GitHub Actions workflows for:

### CI Pipeline (`.github/workflows/ci.yml`)
- Runs on push to main, develop, and claude/* branches
- Backend: Go tests, linting, and build
- Frontend: npm build and linting
- Security: Trivy vulnerability scanning
- Docker: Build test for both services

### Deploy Pipeline (`.github/workflows/deploy.yml`)
- Automated deployment to staging/production
- Docker image building and pushing
- SSH deployment to servers
- Health checks after deployment
- Slack notifications

**Required Secrets**:
- `DOCKER_USERNAME` & `DOCKER_PASSWORD`
- `DEPLOY_HOST`, `DEPLOY_USER`, `DEPLOY_KEY`
- `API_BASE_URL` & `WS_BASE_URL`
- `SLACK_WEBHOOK` (optional)

## Development

### Backend Development

The backend fetches real-time cryptocurrency data from CoinGecko on startup. If the API is unavailable, it falls back to sample data. The `backend/services/coingecko.go` service handles all API interactions with automatic rate limiting.

### Frontend Development

The frontend is configured to connect to the backend at:
- API: `http://localhost:8080/api`
- WebSocket: `ws://localhost:8080/ws`

You can change these URLs in `frontend/nuxt.config.ts` or via environment variables:
- `API_BASE_URL`
- `WS_BASE_URL`

## License

MIT