import { defineConfig, type ServerOptions } from 'vite';
import { readFileSync } from 'node:fs';
import vue from '@vitejs/plugin-vue';

const serverConfig: ServerOptions = {};
if (process.env.ENV === 'local') {
  serverConfig.cors = true;
  serverConfig.https = {
    cert: readFileSync('./serve/cert.pem'),
    key: readFileSync('./serve/key.pem'),
  };
  serverConfig.port = 3000;
}

export default defineConfig({
  plugins: [vue()],
  server: serverConfig,
});
