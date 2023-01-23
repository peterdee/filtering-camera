import { defineConfig } from 'vite';
import { readFileSync } from 'node:fs';
import vue from '@vitejs/plugin-vue';

export default defineConfig({
  plugins: [vue()],
  server: {
    cors: true,
    https: {
      cert: readFileSync('./serve/cert.pem'),
      key: readFileSync('./serve/key.pem'),
    },
    port: 3000,
  },
});
