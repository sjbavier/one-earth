import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";

export default defineConfig({
  plugins: [react()],
  server: { proxy: { "/api": "http://localhost:8080" } },
  define: {
    "import.meta.env.VITE_PUBLIC_SITE_URL": JSON.stringify(
      (import.meta as any).env?.VITE_PUBLIC_SITE_URL ?? ""
    ),
  },
});
