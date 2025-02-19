# Stage 1: Builder Stage - Compile the Go application
# FROM golang:latest AS builder
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Copy go.mod and go.sum first for caching and faster builds
COPY go.mod go.sum ./
# RUN go mod download -x # Download Go modules
RUN go mod tidy

COPY . .

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/main # Or your main file


# Stage 2: Final Stage - Create the minimal runtime image
FROM alpine:latest AS final

WORKDIR /app

# Copy necessary files from the builder stage
COPY --from=builder /app/server /app/server
#COPY --from=builder /app/.env /app/.env # If you have an .env file, copy it
#COPY --from=builder /app/static /app/static # If you have static files, copy them (optional)
# Copy any other required resources (templates, config files etc.)


# Expose the port your API server listens on (adjust if needed, e.g., Fiber default is 3000)
EXPOSE 3000


# Command to run your API server
CMD ["./server"]
