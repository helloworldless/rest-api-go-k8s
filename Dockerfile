FROM golang:1.14

WORKDIR /go/src/app

COPY . /go/src/app

RUN go build -o /go/bin/app /go/src/app/main.go

CMD ["/go/bin/app"]

EXPOSE 8080