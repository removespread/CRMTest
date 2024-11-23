FROM golang:1.23.2-alpine
LABEL maintainer="Viktor Sernyaev <removespread@internet.ru>"
COPY . /app
WORKDIR /app
RUN go mod tidy
EXPOSE 8080
CMD ["go", "run", "main.go"]
