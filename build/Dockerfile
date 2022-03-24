FROM golang:1.17 AS builder

COPY ./as207414.net .
RUN go build web.go

FROM alpine

COPY --from=builder /go/web /app

WORKDIR /root/asn

ENTRYPOINT [ "/app/web" ]
