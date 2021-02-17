FROM golang:alpine

WORKDIR /go/src/app

COPY . .

RUN go build -o ./web_service .

EXPOSE 8080

CMD ["./web_service"]