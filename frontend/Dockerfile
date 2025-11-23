# Build stage
FROM node:20-alpine AS builder

# Set working directory
WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm ci

# Copy source code
COPY . .

# Build arguments for runtime config
ARG API_BASE_URL=http://localhost:8080/api
ARG WS_BASE_URL=ws://localhost:8080/ws

ENV API_BASE_URL=$API_BASE_URL
ENV WS_BASE_URL=$WS_BASE_URL

# Build the application
RUN npm run build

# Runtime stage
FROM node:20-alpine

WORKDIR /app

# Copy built application from builder
COPY --from=builder /app/.output ./.output
COPY --from=builder /app/package*.json ./

# Install only production dependencies
RUN npm ci --only=production

# Expose port
EXPOSE 3000

# Health check
HEALTHCHECK --interval=30s --timeout=3s --start-period=10s --retries=3 \
  CMD wget --no-verbose --tries=1 --spider http://localhost:3000 || exit 1

# Set environment to production
ENV NODE_ENV=production

# Run the application
CMD ["node", ".output/server/index.mjs"]
