// @ts-check
import { defineConfig } from "astro/config";
import aws from "astro-sst";

import tailwindcss from "@tailwindcss/vite";

// https://astro.build/config
export default defineConfig({
  output: "static",
  adapter: aws(),

  vite: {
    plugins: [tailwindcss()],
  },
});
