# Stage 1: Build the Go application
FROM golang:1.23 AS builder  

WORKDIR /app  
# Copy go.mod and go.sum for dependency installation  
COPY go.mod go.sum ./  
 
ENV GOPROXY=https://proxy.golang.org,direct

COPY vendor /app/vendor
ENV GOFLAGS="-mod=vendor"

# Copy the entire application source  
COPY . .  

# Build the application  
RUN CGO_ENABLED=0 GOOS=linux go build -o main ./main.go  

# Stage 2: Create a lightweight container  
FROM alpine:latest  

WORKDIR /app  

# Copy the compiled Go binary from the builder  
COPY --from=builder /app/main .  

EXPOSE 8080  

CMD ["./main"]
