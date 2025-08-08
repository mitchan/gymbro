import { defineConfig } from "vite";
import solid from "vite-plugin-solid";
import tailwindcss from "@tailwindcss/vite";

export default defineConfig({
  plugins: [solid(), tailwindcss()],
  server: {
    host: "0.0.0.0",
    port: 3000,
    hmr: {
      clientPort: 3000,
    },
    proxy: {
      "/api": {
        target: "http://server:8080",
        changeOrigin: true,
        secure: false,
        configure: (proxy) => {
          proxy.on("proxyReq", () => {});
        },
      },
    },
  },
});
