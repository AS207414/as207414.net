FROM alpine:latest as build
RUN apk --update add ca-certificates

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt

ADD as207414_linux_amd64 /

CMD [ "/as207414_linux_amd64" ]