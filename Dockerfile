FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o artaferabackend .

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/artaferabackend .

EXPOSE 8080

CMD ["./artaferabackend"]
