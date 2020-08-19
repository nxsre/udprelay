FROM golang:1.15-alpine AS builder

RUN apk add make git

ADD . /src
WORKDIR /src
RUN make udprelay

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /src/udprelay /usr/local/bin/udprelay

WORKDIR /root
CMD ["/usr/local/bin/udprelay", "9999"]
EXPOSE 9999/udp
