# Build stage
FROM golang:1.20 AS builder
WORKDIR /simulator
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o doroha-simulator ./cmd/

# Final stage
FROM alpine:3.15
WORKDIR /root/
COPY --from=builder /simulator/doroha-simulator .
COPY --from=builder /simulator/filtered_kyiv.geojson .
CMD ["./doroha-simulator"]