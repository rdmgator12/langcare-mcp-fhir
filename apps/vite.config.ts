import { defineConfig } from "vite";
import react from "@vitejs/plugin-react";
import { viteSingleFile } from "vite-plugin-singlefile";
import path from "node:path";

const INPUT = process.env.INPUT;
if (!INPUT) {
  throw new Error("INPUT env var required (e.g. INPUT=fhir-explorer/index.html)");
}

const OUTDIR = process.env.OUTDIR || "dist";
const isDev = process.env.NODE_ENV === "development";

// Derive root from INPUT so output is flat (no nested dirs)
const appRoot = path.dirname(INPUT);

export default defineConfig({
  root: appRoot,
  plugins: [react(), viteSingleFile()],
  build: {
    sourcemap: isDev ? "inline" : undefined,
    cssMinify: !isDev,
    minify: !isDev,
    rollupOptions: { input: path.resolve(INPUT) },
    outDir: path.resolve(OUTDIR),
    emptyOutDir: true,
  },
});
