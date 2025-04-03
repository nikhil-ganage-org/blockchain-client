# Build stage
FROM golang:1.19 AS builder
WORKDIR /app

COPY *.go ./
COPY go.mod* go.sum* ./

RUN if [ ! -f go.mod ]; then \
      go mod init blockchain-client && \
      go mod tidy; \
    fi

RUN go mod download
RUN go test -v
RUN CGO_ENABLED=0 GOOS=linux go build -o blockchain-client

# Final stage with SSL certificates
FROM alpine:latest
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY --from=builder /app/blockchain-client .
EXPOSE 8080
CMD ["./blockchain-client"]