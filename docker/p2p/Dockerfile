FROM golang:1.16.3-alpine as builder

RUN apk update \
    && apk upgrade \
    && apk add --no-cache make

WORKDIR /src/challenge
COPY ../.. .
RUN make compile/p2p

FROM alpine:3.11

WORKDIR /challenge

COPY --from=builder /src/challenge/build/p2p /usr/bin/challenge

CMD ["challenge"]
