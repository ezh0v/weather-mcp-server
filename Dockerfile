FROM golang:1.24-alpine3.21 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o weather-mcp-server ./cmd/weather-mcp-server

FROM gcr.io/distroless/base-debian12

WORKDIR /app

COPY --from=builder /app/weather-mcp-server .

USER nonroot:nonroot

EXPOSE 8000

ENTRYPOINT ["./weather-mcp-server", "--address", "0.0.0.0:8000"]