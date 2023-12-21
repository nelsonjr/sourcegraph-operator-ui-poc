import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  server: {
    port: 8889,
    proxy: {
      "/api": {
        target: "http://localhost:8888",
      },
    },
  },
  plugins: [react()],
  define: {
    "process.env.API_ENDPOINT": JSON.stringify(""),
  },
});
