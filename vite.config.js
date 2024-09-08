import { defineConfig } from "vite";
import react from "@vitejs/plugin-react-swc";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  server: {
    port: 5000,
    proxy: {
      "/api": {
        target: "http://0.0.0.0:3000/",
        changeOrigin: true,
        secure: false,
      },
    },
  },
});
