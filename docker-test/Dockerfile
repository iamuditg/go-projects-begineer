# use an offcial golang runtime as a parent image
FROM golang:latest

RUN mkdir /app

ADD . /app

WORKDIR  /app

RUN go build -o main .

CMD ["/app/main"]
