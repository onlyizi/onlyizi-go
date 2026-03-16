FROM golang:1.25.5

WORKDIR /app

COPY . .

RUN go build -o example cmd/example.go

CMD ["./example"]