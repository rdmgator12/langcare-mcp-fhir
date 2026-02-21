#!/usr/bin/env bash
set -euo pipefail

APPS_DIR="apps"
EMBED_DIR="internal/apps/dist"

cd "$APPS_DIR"
npm install --prefer-offline

# Clean previous build
rm -rf dist dist-tmp

# Build each app into its own temp dir, then collect into dist/
for app_dir in */; do
  app_name="${app_dir%/}"
  if [ -f "$app_dir/index.html" ]; then
    echo "Building app: $app_name"
    INPUT="$app_dir/index.html" OUTDIR="dist-tmp/$app_name" npx vite build
    mkdir -p dist
    mv "dist-tmp/$app_name/index.html" "dist/$app_name.html"
  fi
done

# Clean up temp dirs
rm -rf dist-tmp

cd ..

# Copy built bundles to Go embed directory
mkdir -p "$EMBED_DIR"
cp "$APPS_DIR"/dist/*.html "$EMBED_DIR/"

echo "Apps built and copied to $EMBED_DIR:"
ls -la "$EMBED_DIR"
