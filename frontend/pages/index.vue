<template>
  <div class="container">
    <header class="header">
      <h1>Crypto Marketplace</h1>
      <div class="connection-status">
        <span class="status-indicator" :class="{ connected: wsConnected }"></span>
        <span>{{ wsConnected ? 'Live' : 'Disconnected' }}</span>
      </div>
    </header>

    <div class="products-grid">
      <ProductCard
        v-for="product in products"
        :key="product.id"
        :product="product"
        @update="handleUpdate"
        @delete="handleDelete"
      />
    </div>

    <div v-if="products.length === 0" class="empty-state">
      <p>No products available</p>
    </div>

    <button class="add-button" @click="showAddForm = true">
      + Add Product
    </button>

    <AddProductForm
      v-if="showAddForm"
      @close="showAddForm = false"
      @add="handleAdd"
    />
  </div>
</template>

<script setup lang="ts">
import type { Product, ProductUpdate } from '~/composables/useApi'

const { getProducts } = useApi()
const { connect, disconnect, isConnected: wsConnected } = useWebSocket()

const products = ref<Product[]>([])
const showAddForm = ref(false)

// Load initial products
onMounted(async () => {
  products.value = await getProducts()

  // Connect to WebSocket for real-time updates
  connect(handleWebSocketUpdate)
})

onBeforeUnmount(() => {
  disconnect()
})

const handleWebSocketUpdate = (update: ProductUpdate) => {
  console.log('Received update:', update)

  switch (update.type) {
    case 'new_product':
      // Add new product if it doesn't exist
      if (!products.value.find(p => p.id === update.product.id)) {
        products.value.push(update.product)
      }
      break

    case 'price_update':
      // Update existing product
      const index = products.value.findIndex(p => p.id === update.product.id)
      if (index !== -1) {
        products.value[index] = update.product
      }
      break

    case 'delete_product':
      // Remove deleted product
      products.value = products.value.filter(p => p.id !== update.product.id)
      break
  }
}

const handleUpdate = (updatedProduct: Product) => {
  const index = products.value.findIndex(p => p.id === updatedProduct.id)
  if (index !== -1) {
    products.value[index] = updatedProduct
  }
}

const handleDelete = (id: string) => {
  products.value = products.value.filter(p => p.id !== id)
}

const handleAdd = (newProduct: Product) => {
  products.value.push(newProduct)
  showAddForm.value = false
}
</script>

<style scoped>
.container {
  max-width: 1400px;
  margin: 0 auto;
  padding: 2rem;
}

.header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
  padding-bottom: 1rem;
  border-bottom: 2px solid #1e1e3f;
}

.header h1 {
  font-size: 2.5rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.connection-status {
  display: flex;
  align-items: center;
  gap: 0.5rem;
  padding: 0.5rem 1rem;
  background: #1e1e3f;
  border-radius: 20px;
}

.status-indicator {
  width: 10px;
  height: 10px;
  border-radius: 50%;
  background: #dc3545;
  animation: pulse 2s infinite;
}

.status-indicator.connected {
  background: #28a745;
}

@keyframes pulse {
  0%, 100% { opacity: 1; }
  50% { opacity: 0.5; }
}

.products-grid {
  display: grid;
  grid-template-columns: repeat(auto-fill, minmax(350px, 1fr));
  gap: 1.5rem;
  margin-bottom: 2rem;
}

.empty-state {
  text-align: center;
  padding: 4rem;
  color: #666;
}

.add-button {
  position: fixed;
  bottom: 2rem;
  right: 2rem;
  padding: 1rem 2rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
  border: none;
  border-radius: 50px;
  font-size: 1.1rem;
  font-weight: 600;
  cursor: pointer;
  box-shadow: 0 4px 15px rgba(102, 126, 234, 0.4);
  transition: transform 0.2s, box-shadow 0.2s;
}

.add-button:hover {
  transform: translateY(-2px);
  box-shadow: 0 6px 20px rgba(102, 126, 234, 0.6);
}

.add-button:active {
  transform: translateY(0);
}
</style>
