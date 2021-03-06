FROM alpine:latest as dumb-init
RUN apk add --no-cache build-base git bash
RUN git clone https://github.com/Yelp/dumb-init.git
WORKDIR /dumb-init
RUN make

FROM golang:1.13-alpine as builder
RUN apk --no-cache add make git gcc libtool musl-dev upx
ENV GO111MODULE=on
COPY . /platform
WORKDIR /platform
RUN make build

FROM alpine:latest
COPY --from=dumb-init /dumb-init/dumb-init /bin/dumb-init
RUN apk add ca-certificates && \
    rm -rf /var/cache/apk/* /tmp/* && \
    [ ! -e /etc/nsswitch.conf ] && echo 'hosts: files dns' > /etc/nsswitch.conf

WORKDIR /
COPY . /
COPY --from=builder /platform/platform /platform
COPY entrypoint.sh /
RUN chmod 755 entrypoint.sh
ENTRYPOINT ["./entrypoint.sh"]
