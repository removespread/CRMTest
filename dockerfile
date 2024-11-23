FROM golang:1.23.2-alpine
LABEL maintainer="Viktor Sernyaev <removespread@internet.ru>"
COPY go.mod go.sum ./
RUN go mod download
WORKDIR /app
COPY . .
RUN go build -o app ./cmd/main.go 
EXPOSE 8080
COPY ./builder/entrypoint.sh /entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
CMD ["./app"]
