import { sveltekit } from '@sveltejs/kit/vite';
import { defineConfig } from 'vite';


/** @type {import('vite').UserConfig} */
const config = {
  plugins: [sveltekit()],
  server: {
    proxy: {
      // Proxy API requests from /api/* to localhost:3000
      '/api': {
        target: 'http://localhost:3000',
        changeOrigin: true,
        rewrite: path => path.replace(/^\/api/, '') // optional: strip '/api' before sending to Go
      }
    }
  }
};

export default config;


