import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { viteStaticCopy } from "vite-plugin-static-copy";

export default defineConfig({
  plugins: [
    react(),
    viteStaticCopy({
      targets: [
        // Adjust paths if your front-end is not at repo root
        { src: "openapi/openapi.yaml", dest: "/openapi.yaml" }, // => /openapi.yaml
      ],
    }),
  ],
  server: { proxy: { "/api": "http://localhost:8080" } },
  define: {
    "import.meta.env.VITE_PUBLIC_SITE_URL": JSON.stringify(
      (import.meta as any).env?.VITE_PUBLIC_SITE_URL ?? ""
    ),
  },
});
