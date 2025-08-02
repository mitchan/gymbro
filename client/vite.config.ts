import { defineConfig } from "vite";
import solid from "vite-plugin-solid";

export default defineConfig({
  plugins: [solid()],
  server: {
    // allowedHosts: ["localhost"],
    host: "0.0.0.0",
    port: 3000,
    hmr: {
      clientPort: 3000,
    },
    // watch: {
    //   usePolling: true,
    // },
  },
});
