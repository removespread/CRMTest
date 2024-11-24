FROM golang:1.23.2-alpine AS builder
LABEL maintainer="Viktor Sernyaev <removespread@internet.ru>"
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o app ./cmd/main.go
EXPOSE 8080
CMD ["./app"]