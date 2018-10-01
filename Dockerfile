FROM golang

RUN rm -rf /go/src/github.com/spotted/auth
COPY . /go/src/github.com/spotted/auth
WORKDIR /go/src/github.com/spotted/auth

RUN ["/bin/bash", "-c", "go get -v -d ./..."]

RUN CGO_ENABLED=0 GOOS=linux go build -o /bin/auth .

FROM alpine
RUN apk add --no-cache bash
RUN apk add --no-cache ca-certificates

EXPOSE 8080

CMD /bin/auth