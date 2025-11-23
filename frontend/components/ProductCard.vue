<template>
  <div class="product-card">
    <div class="card-header">
      <div class="symbol-badge">{{ product.symbol }}</div>
      <div class="actions">
        <button @click="toggleEdit" class="btn-icon" title="Edit">✎</button>
        <button @click="confirmDelete" class="btn-icon delete" title="Delete">×</button>
      </div>
    </div>

    <h2 class="product-name">{{ product.name }}</h2>

    <div class="price-section">
      <div class="price">${{ formatPrice(product.price) }}</div>
      <div class="change" :class="changeClass">
        {{ product.change24h > 0 ? '↑' : '↓' }} {{ Math.abs(product.change24h) }}%
      </div>
    </div>

    <div class="stats">
      <div class="stat">
        <span class="stat-label">24h Volume</span>
        <span class="stat-value">${{ formatLargeNumber(product.volume24h) }}</span>
      </div>
      <div class="stat">
        <span class="stat-label">Market Cap</span>
        <span class="stat-value">${{ formatLargeNumber(product.marketCap) }}</span>
      </div>
    </div>

    <p class="description">{{ product.description }}</p>

    <div v-if="isEditing" class="edit-form">
      <input v-model="editForm.name" placeholder="Name" class="input" />
      <input v-model="editForm.symbol" placeholder="Symbol" class="input" />
      <input v-model.number="editForm.price" type="number" step="0.01" placeholder="Price" class="input" />
      <input v-model.number="editForm.change24h" type="number" step="0.01" placeholder="24h Change %" class="input" />
      <input v-model.number="editForm.volume24h" type="number" placeholder="24h Volume" class="input" />
      <input v-model.number="editForm.marketCap" type="number" placeholder="Market Cap" class="input" />
      <textarea v-model="editForm.description" placeholder="Description" class="input textarea"></textarea>

      <div class="form-actions">
        <button @click="saveEdit" class="btn btn-primary">Save</button>
        <button @click="cancelEdit" class="btn btn-secondary">Cancel</button>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import type { Product } from '~/composables/useApi'

const props = defineProps<{
  product: Product
}>()

const emit = defineEmits<{
  update: [product: Product]
  delete: [id: string]
}>()

const { updateProduct, deleteProduct } = useApi()

const isEditing = ref(false)
const editForm = ref({ ...props.product })

const changeClass = computed(() => ({
  positive: props.product.change24h > 0,
  negative: props.product.change24h < 0
}))

const formatPrice = (price: number) => {
  return price.toLocaleString('en-US', {
    minimumFractionDigits: 2,
    maximumFractionDigits: 2
  })
}

const formatLargeNumber = (num: number) => {
  if (num >= 1e9) return (num / 1e9).toFixed(2) + 'B'
  if (num >= 1e6) return (num / 1e6).toFixed(2) + 'M'
  if (num >= 1e3) return (num / 1e3).toFixed(2) + 'K'
  return num.toString()
}

const toggleEdit = () => {
  isEditing.value = !isEditing.value
  if (isEditing.value) {
    editForm.value = { ...props.product }
  }
}

const saveEdit = async () => {
  const updated = await updateProduct(props.product.id, editForm.value)
  if (updated) {
    emit('update', updated)
    isEditing.value = false
  }
}

const cancelEdit = () => {
  isEditing.value = false
  editForm.value = { ...props.product }
}

const confirmDelete = async () => {
  if (confirm(`Are you sure you want to delete ${props.product.name}?`)) {
    const success = await deleteProduct(props.product.id)
    if (success) {
      emit('delete', props.product.id)
    }
  }
}
</script>

<style scoped>
.product-card {
  background: linear-gradient(135deg, #1e1e3f 0%, #2a2a4f 100%);
  border-radius: 16px;
  padding: 1.5rem;
  border: 1px solid #3a3a5f;
  transition: transform 0.2s, box-shadow 0.2s;
}

.product-card:hover {
  transform: translateY(-4px);
  box-shadow: 0 8px 24px rgba(102, 126, 234, 0.2);
}

.card-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1rem;
}

.symbol-badge {
  background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
  padding: 0.5rem 1rem;
  border-radius: 8px;
  font-weight: 700;
  font-size: 0.9rem;
}

.actions {
  display: flex;
  gap: 0.5rem;
}

.btn-icon {
  background: #2a2a4f;
  border: 1px solid #3a3a5f;
  color: #fff;
  width: 32px;
  height: 32px;
  border-radius: 8px;
  cursor: pointer;
  font-size: 1.2rem;
  transition: all 0.2s;
}

.btn-icon:hover {
  background: #3a3a5f;
  transform: scale(1.1);
}

.btn-icon.delete:hover {
  background: #dc3545;
  border-color: #dc3545;
}

.product-name {
  font-size: 1.8rem;
  margin-bottom: 1rem;
  color: #fff;
}

.price-section {
  display: flex;
  align-items: baseline;
  gap: 1rem;
  margin-bottom: 1.5rem;
}

.price {
  font-size: 2rem;
  font-weight: 700;
  color: #fff;
}

.change {
  font-size: 1.1rem;
  font-weight: 600;
  padding: 0.25rem 0.75rem;
  border-radius: 6px;
}

.change.positive {
  background: rgba(40, 167, 69, 0.2);
  color: #28a745;
}

.change.negative {
  background: rgba(220, 53, 69, 0.2);
  color: #dc3545;
}

.stats {
  display: grid;
  grid-template-columns: 1fr 1fr;
  gap: 1rem;
  margin-bottom: 1rem;
  padding: 1rem;
  background: rgba(0, 0, 0, 0.2);
  border-radius: 8px;
}

.stat {
  display: flex;
  flex-direction: column;
  gap: 0.25rem;
}

.stat-label {
  font-size: 0.8rem;
  color: #9999cc;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.stat-value {
  font-size: 1.1rem;
  font-weight: 600;
  color: #fff;
}

.description {
  color: #ccccdd;
  line-height: 1.6;
  margin-bottom: 1rem;
}

.edit-form {
  margin-top: 1.5rem;
  padding-top: 1.5rem;
  border-top: 1px solid #3a3a5f;
}

.input {
  width: 100%;
  padding: 0.75rem;
  margin-bottom: 0.75rem;
  background: #2a2a4f;
  border: 1px solid #3a3a5f;
  border-radius: 8px;
  color: #fff;
  font-size: 1rem;
}

.input:focus {
  outline: none;
  border-color: #667eea;
}

.textarea {
  min-height: 80px;
  resize: vertical;
}

.form-actions {
  display: flex;
  gap: 0.75rem;
}

.btn {
  flex: 1;
  padding: 0.75rem;
  border: none;
  border-radius: 8px;
  font-weight: 600;
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
