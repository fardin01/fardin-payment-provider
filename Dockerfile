FROM golang:1.16-alpine AS builder


WORKDIR /app
RUN apk --no-cache add make git

COPY Makefile go.mod go.sum ./
RUN make mod tools

COPY . .
RUN make lint build

FROM alpine

MAINTAINER Fardin Khanjani <khanjani.fardin@gmail.com>

COPY --from=builder /app/bin/fardin-payment-provider /fardin-payment-provider

EXPOSE 9000

CMD ["/fardin-payment-provider"]
