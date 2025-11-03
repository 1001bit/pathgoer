FROM golang:1.25.3 AS builder

WORKDIR /app

# dependencies
COPY go.* .
RUN go mod download
# other files
COPY . .

# compiling go -> ./main
RUN go build -o main main.go

# Final stage
FROM debian:bookworm-slim AS runner

WORKDIR /app

COPY --from=builder /app/main ./main

ENTRYPOINT ["./main"]