FROM golang:1.25-alpine

# Install CGO dependencies for SQLite
RUN apk add --no-cache gcc musl-dev

WORKDIR /app

COPY go.* ./
RUN go mod download

COPY . .

# Build with CGO enabled
RUN CGO_ENABLED=1 go build -o main ./cmd/server

EXPOSE 8080

CMD ["./main"]