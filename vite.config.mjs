import { defineConfig } from 'vite';
import react from '@vitejs/plugin-react-swc';

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5000,
    proxy: {
      'go/api': {
        target: 'http://0.0.0.0:4000/',
        changeOrigin: true,
        secure: false,
      },
      'nest/api': {
        target: 'http://0.0.0.0:3000/',
        changeOrigin: true,
        secure: false,
      },
    },
  },
  build: {
    outDir: 'dist/client', // Specify your desired output directory here
  },
});
