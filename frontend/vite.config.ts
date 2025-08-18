import { defineConfig } from 'vite'
import react from '@vitejs/plugin-react'

export default defineConfig({
  plugins: [react()],
  server: { proxy: { '/api': 'http://localhost:8081' } },
  define: {
    'import.meta.env.VITE_PUBLIC_SITE_URL': JSON.stringify(process.env.VITE_PUBLIC_SITE_URL || '')
  }
})
