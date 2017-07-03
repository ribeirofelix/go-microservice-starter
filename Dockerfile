FROM golang:latest
WORKDIR /go/src/github.com/hellomd/go-microservice-starter
ADD . .
RUN go install

ENV PORT 3000
ENV APP_NAME go-microservice-starter

EXPOSE 3000
ENTRYPOINT /go/bin/go-microservice-starter
