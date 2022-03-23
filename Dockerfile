FROM golang:1.17 AS builder

COPY ./as207414.net .
RUN go build main.go

FROM alpine

COPY --from=builder /go/web /root/asn

WORKDIR /root/asn

ENTRYPOINT [ "/main" ]
