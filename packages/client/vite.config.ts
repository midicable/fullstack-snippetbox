import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

// https://vitejs.dev/config/
export default defineConfig({
  plugins: [react()],
  resolve: {
    alias: {
      app: "/src/app",
      pages: "/src/pages",
      entities: "/src/entities",
      widgets: "/src/widgets",
      shared: "/src/shared",
    },
  },
});
