FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o mcp-brewfather ./cmd/mcp-brewfather

FROM gcr.io/distroless/static:nonroot
COPY --from=builder /app/mcp-brewfather /mcp-brewfather
ENTRYPOINT ["/mcp-brewfather"]
