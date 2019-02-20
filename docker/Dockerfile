FROM alpine:3.9

WORKDIR /go/bin
ADD scrape .
RUN apk update && apk add openssl ca-certificates

ENTRYPOINT /go/bin/scrape


