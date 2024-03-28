# Start from the official Go image
FROM golang:1.20
# Install Node.js, npm, and Java
RUN apt-get update && \
    apt-get install -y nodejs npm && \
    apt-get clean && \
    npm install -g @openapitools/openapi-generator-cli && \
    apt-get install -y default-jdk && \
    apt-get clean
# Set the Current Working Directory inside the container.
WORKDIR /api
# Copy go mod and sum files
COPY go.mod go.sum ./
# Download all dependencies.
RUN go mod download
# Copy the project directory (filtered by .dockerignore)
COPY . .
# Build your application.
RUN CGO_ENABLED=0 GOOS=linux go generate ./cmd/
RUN CGO_ENABLED=0 GOOS=linux go build -o doroha-api ./cmd/
# Run
CMD ["./doroha-api"]
