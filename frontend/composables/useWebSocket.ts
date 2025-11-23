import type { ProductUpdate } from './useApi'

export const useWebSocket = () => {
  const config = useRuntimeConfig()
  const wsBase = config.public.wsBase

  let ws: WebSocket | null = null
  let reconnectTimeout: NodeJS.Timeout | null = null
  const isConnected = ref(false)
  const reconnectAttempts = ref(0)
  const maxReconnectAttempts = 5

  const connect = (onMessage: (update: ProductUpdate) => void) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      console.log('WebSocket already connected')
      return
    }

    try {
      ws = new WebSocket(wsBase)

      ws.onopen = () => {
        console.log('WebSocket connected')
        isConnected.value = true
        reconnectAttempts.value = 0
      }

      ws.onmessage = (event) => {
        try {
          const update: ProductUpdate = JSON.parse(event.data)
          onMessage(update)
        } catch (error) {
          console.error('Error parsing WebSocket message:', error)
        }
      }

      ws.onerror = (error) => {
        console.error('WebSocket error:', error)
      }

      ws.onclose = () => {
        console.log('WebSocket disconnected')
        isConnected.value = false

        // Attempt to reconnect
        if (reconnectAttempts.value < maxReconnectAttempts) {
          reconnectAttempts.value++
          const delay = Math.min(1000 * Math.pow(2, reconnectAttempts.value), 30000)
          console.log(`Reconnecting in ${delay}ms... (attempt ${reconnectAttempts.value}/${maxReconnectAttempts})`)

          reconnectTimeout = setTimeout(() => {
            connect(onMessage)
          }, delay)
        } else {
          console.error('Max reconnection attempts reached')
        }
      }
    } catch (error) {
      console.error('Error creating WebSocket:', error)
    }
  }

  const disconnect = () => {
    if (reconnectTimeout) {
      clearTimeout(reconnectTimeout)
      reconnectTimeout = null
    }

    if (ws) {
      ws.close()
      ws = null
    }

    isConnected.value = false
    reconnectAttempts.value = 0
  }

  const send = (data: any) => {
    if (ws && ws.readyState === WebSocket.OPEN) {
      ws.send(JSON.stringify(data))
    } else {
      console.error('WebSocket is not connected')
    }
  }

  return {
    connect,
    disconnect,
    send,
    isConnected: readonly(isConnected)
  }
}
