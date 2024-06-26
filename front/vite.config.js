import { defineConfig } from 'vite'
import vue from '@vitejs/plugin-vue'
import { resolve } from 'path'

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [vue()],
  envPrefix: "APP_",
  server: {
    port: 3000,
    open: true
  },

  resolve: {
    alias: {
      '@': resolve(__dirname, 'src'),
      'api': '@/api',
      'components': '@/stuView',
      'assets': '@/assets'
    }
  }
})
