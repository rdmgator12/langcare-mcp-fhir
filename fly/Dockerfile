# ── Stage 0: Build MCP App Bundles (Node) ──
FROM node:20-alpine AS app-builder

WORKDIR /apps
COPY apps/package.json apps/package-lock.json* ./
RUN npm install --prefer-offline

COPY apps/ ./

# Build each app into a per-app temp dir, then collect
RUN for dir in */; do \
      name="${dir%/}"; \
      if [ -f "$dir/index.html" ]; then \
        INPUT="$dir/index.html" OUTDIR="dist-tmp/$name" npx vite build && \
        mkdir -p dist && \
        mv "dist-tmp/$name/index.html" "dist/$name.html"; \
      fi; \
    done && rm -rf dist-tmp

# ── Stage 1: Build Go Binary ──
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

# Copy built app bundles into the Go embed directory
COPY --from=app-builder /apps/dist/*.html /src/internal/apps/dist/

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o /mcp-server ./cmd/server

# ── Stage 2: Runtime ──
FROM alpine:3.21

RUN apk add --no-cache ca-certificates
RUN adduser -D -u 1001 mcpuser

WORKDIR /app
COPY --from=builder /mcp-server /app/mcp-server
COPY fly/config.fly.epic.yaml /app/config.fly.epic.yaml
COPY fly/config.fly.gcp.yaml /app/config.fly.gcp.yaml
COPY fly/config.fly.cerner.yaml /app/config.fly.cerner.yaml
COPY fly/docker-entrypoint.sh /app/docker-entrypoint.sh

RUN chmod +x /app/docker-entrypoint.sh

USER mcpuser
EXPOSE 8080

HEALTHCHECK --interval=30s --timeout=3s --start-period=5s \
    CMD wget -qO- http://localhost:8080/health || exit 1

ENTRYPOINT ["/app/docker-entrypoint.sh"]
