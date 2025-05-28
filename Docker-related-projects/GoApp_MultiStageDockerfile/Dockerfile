# -------- Stage 1: Build --------
    FROM golang:1.24-alpine AS builder
    WORKDIR /app
    COPY . .
    RUN go build -o main .
    # ----- Stage 2: Runtime --------
    FROM alpine:latest
    WORKDIR /root/
    COPY --from=builder /app/main .
    CMD ["./main"]