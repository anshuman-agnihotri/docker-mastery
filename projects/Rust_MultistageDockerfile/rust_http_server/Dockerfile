# Stage 1: Build
FROM rust:1.75 AS builder
WORKDIR /app
COPY . .
RUN cargo build --release
# Stage 2: Runtime with same OS
FROM rust:1.75-slim AS runtime
WORKDIR /app
COPY --from=builder /app/target/release/rust_http_server .
EXPOSE 3000
CMD ["./rust_http_server"]