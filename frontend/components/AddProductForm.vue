<template>
  <div class="modal-overlay" @click.self="emit('close')">
    <div class="modal">
      <div class="modal-header">
        <h2>Add New Product</h2>
        <button class="close-btn" @click="emit('close')">×</button>
      </div>

      <form @submit.prevent="handleSubmit" class="form">
        <div class="form-group">
          <label>Name</label>
          <input v-model="form.name" required class="input" placeholder="e.g., Bitcoin" />
        </div>

        <div class="form-group">
          <label>Symbol</label>
          <input v-model="form.symbol" required class="input" placeholder="e.g., BTC" />
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>Price ($)</label>
            <input v-model.number="form.price" type="number" step="0.01" required class="input" placeholder="0.00" />
          </div>

          <div class="form-group">
            <label>24h Change (%)</label>
            <input v-model.number="form.change24h" type="number" step="0.01" required class="input" placeholder="0.00" />
          </div>
        </div>

        <div class="form-row">
          <div class="form-group">
            <label>24h Volume ($)</label>
            <input v-model.number="form.volume24h" type="number" required class="input" placeholder="0" />
          </div>

          <div class="form-group">
            <label>Market Cap ($)</label>
            <input v-model.number="form.marketCap" type="number" required class="input" placeholder="0" />
          </div>
        </div>

        <div class="form-group">
          <label>Description</label>
          <textarea v-model="form.description" required class="input textarea" placeholder="Brief description..."></textarea>
        </div>

        <div class="form-actions">
          <button type="button" @click="emit('close')" class="btn btn-secondary">Cancel</button>
          <button type="submit" class="btn btn-primary">Add Product</button>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Product } from '~/composables/useApi'

const emit = defineEmits<{
  close: []
  add: [product: Product]
}>()

const { createProduct } = useApi()

const form = ref({
  name: '',
  symbol: '',
  price: 0,
  change24h: 0,
  volume24h: 0,
  marketCap: 0,
  description: ''
})

const handleSubmit = async () => {
  const newProduct = await createProduct(form.value)
  if (newProduct) {
    emit('add', newProduct)
  }
}
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  bottom: 0;
  background: rgba(0, 0, 0, 0.7);
  display: flex;
  align-items: center;
  justify-content: center;
  z-index: 1000;
  backdrop-filter: blur(4px);
}

.modal {
  background: linear-gradient(135deg, #1e1e3f 0%, #2a2a4f 100%);
  border-radius: 16px;
  padding: 2rem;
  max-width: 600px;
  width: 90%;
  max-height: 90vh;
  overflow-y: auto;
  border: 1px solid #3a3a5f;
  box-shadow: 0 20px 60px rgba(0, 0, 0, 0.5);
}

.modal-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 2rem;
}

.modal-header h2 {
  font-size: 1.8rem;
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  -webkit-background-clip: text;
  -webkit-text-fill-color: transparent;
  background-clip: text;
}

.close-btn {
  background: none;
  border: none;
  color: #fff;
  font-size: 2rem;
  cursor: pointer;
  width: 40px;
  height: 40px;
  display: flex;
  align-items: center;
  justify-content: center;
  border-radius: 8px;
  transition: background 0.2s;
}

.close-btn:hover {
  background: rgba(255, 255, 255, 0.1);
}

.form {
  display: flex;
  flex-direction: column;
  gap: 1.25rem;
}

.form-group {
  display: flex;
  flex-direction: column;
  gap: 0.5rem;
}

.form-row {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
}

label {
  font-size: 0.9rem;
  font-weight: 600;
  color: #9999cc;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.input {
  width: 100%;
  padding: 0.875rem;
  background: #2a2a4f;
  border: 1px solid #3a3a5f;
  border-radius: 8px;
  color: #fff;
  font-size: 1rem;
  transition: border-color 0.2s;
}

.input:focus {
  outline: none;
  border-color: #667eea;
}

.textarea {
  min-height: 100px;
  resize: vertical;
  font-family: inherit;
}

.form-actions {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-top: 1rem;
}

.btn {
  padding: 1rem;
  border: none;
  border-radius: 8px;
  font-weight: 600;
  font-size: 1rem;
  cursor: pointer;
  transition: all 0.2s;
}

.btn-primary {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  color: white;
}

.btn-primary:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 12px rgba(102, 126, 234, 0.4);
}

.btn-secondary {
  background: #3a3a5f;
  color: white;
}

.btn-secondary:hover {
  background: #4a4a6f;
}
</style>
