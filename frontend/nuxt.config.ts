// https://nuxt.com/docs/api/configuration/nuxt-config
export default defineNuxtConfig({
  devtools: { enabled: true },

  app: {
    head: {
      title: 'Crypto Marketplace',
      meta: [
        { charset: 'utf-8' },
        { name: 'viewport', content: 'width=device-width, initial-scale=1' },
        { name: 'description', content: 'Real-time cryptocurrency marketplace' }
      ]
    }
  },

  runtimeConfig: {
    public: {
      apiBase: process.env.API_BASE_URL || 'http://localhost:8080/api',
      wsBase: process.env.WS_BASE_URL || 'ws://localhost:8080/ws'
    }
  },

  css: ['~/assets/css/main.css'],

  compatibilityDate: '2024-01-01'
})
