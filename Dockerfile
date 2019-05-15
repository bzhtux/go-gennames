FROM golang:alpine as builder

LABEL maintainer "Yannick Foeillet <bzhtux@gmail.com>"

RUN mkdir -p $GOPATH/src/github.com/bzhtux/go-gennames \
    && apk add git

COPY ./main.go $GOPATH/src/github.com/bzhtux/go-gennames

WORKDIR $GOPATH/src/github.com/bzhtux/go-gennames

RUN go get ./... \
    && go build -o /tmp/gen_names . \
    && ls -ltr /tmp/

FROM alpine
RUN adduser -S -D -H -h /app appuser
USER appuser
COPY --from=builder /tmp/gen_names /usr/bin/
# WORKDIR /app
ENTRYPOINT ["/usr/bin/gen_names"]
