FROM golang

RUN rm -rf /go/src/github.com/x64puzzle/spotted-auth
COPY . /go/src/github.com/x64puzzle/spotted-auth
WORKDIR /go/src/github.com/x64puzzle/spotted-auth

RUN ["/bin/bash", "-c", "go get -v -d ./..."]

RUN go build .

EXPOSE 8888

CMD ./spotted-auth