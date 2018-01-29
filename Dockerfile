FROM golang

ADD . /go/src/github.com/amandhora/cache-proxy

RUN go get -v github.com/amandhora/cache-proxy
RUN go install github.com/amandhora/cache-proxy

ENTRYPOINT /go/bin/cache-proxy

EXPOSE 8080
