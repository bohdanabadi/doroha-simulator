# Build stage
FROM golang:1.20 AS builder
WORKDIR /api
COPY go.mod go.sum ./
RUN go mod download
COPY . .
# Install Node.js, npm, and necessary packages
RUN apt-get update && \
    apt-get install -y nodejs npm && \
    npm install -g @openapitools/openapi-generator-cli && \
    apt-get install -y default-jdk && \
    apt-get clean
RUN CGO_ENABLED=0 GOOS=linux go generate ./cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -o doroha-api ./cmd/

# Final stage
FROM alpine:3.15
WORKDIR /root/
COPY --from=builder /api/doroha-api .
COPY --from=builder /api/config.yml .
CMD ["./doroha-api"]