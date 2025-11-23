export interface Product {
  id: string
  name: string
  symbol: string
  price: number
  change24h: number
  volume24h: number
  marketCap: number
  description: string
  createdAt: string
  updatedAt: string
}

export interface ProductUpdate {
  type: 'price_update' | 'new_product' | 'delete_product'
  product: Product
}

export const useApi = () => {
  const config = useRuntimeConfig()
  const apiBase = config.public.apiBase

  const getProducts = async (): Promise<Product[]> => {
    try {
      const response = await $fetch<Product[]>(`${apiBase}/products`)
      return response
    } catch (error) {
      console.error('Error fetching products:', error)
      return []
    }
  }

  const getProduct = async (id: string): Promise<Product | null> => {
    try {
      const response = await $fetch<Product>(`${apiBase}/products/${id}`)
      return response
    } catch (error) {
      console.error('Error fetching product:', error)
      return null
    }
  }

  const createProduct = async (product: Partial<Product>): Promise<Product | null> => {
    try {
      const response = await $fetch<Product>(`${apiBase}/products`, {
        method: 'POST',
        body: product
      })
      return response
    } catch (error) {
      console.error('Error creating product:', error)
      return null
    }
  }

  const updateProduct = async (id: string, product: Partial<Product>): Promise<Product | null> => {
    try {
      const response = await $fetch<Product>(`${apiBase}/products/${id}`, {
        method: 'PUT',
        body: product
      })
      return response
    } catch (error) {
      console.error('Error updating product:', error)
      return null
    }
  }

  const deleteProduct = async (id: string): Promise<boolean> => {
    try {
      await $fetch(`${apiBase}/products/${id}`, {
        method: 'DELETE'
      })
      return true
    } catch (error) {
      console.error('Error deleting product:', error)
      return false
    }
  }

  return {
    getProducts,
    getProduct,
    createProduct,
    updateProduct,
    deleteProduct
  }
}
