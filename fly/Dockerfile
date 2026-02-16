# ── Stage 1: Build ──
FROM golang:1.25-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .

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
