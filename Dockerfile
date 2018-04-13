FROM golang:latest

RUN mkdir /go/src/golang_webapidemo
ADD . /go/src/golang_webapidemo
WORKDIR /go/src/golang_webapidemo

RUN go get ./...
RUN go build -o main .

CMD ["/go/src/golang_webapidemo/","main"]