# FROM alpine:3.19
FROM alpine@sha256:c5b1261d6d3e43071626931fc004f70149baeba2c8ec672bd4f27761f8e1ad6b

RUN apk add --no-cache socat

WORKDIR /
COPY game.sh flag.txt ./

RUN adduser -DH ctf
USER ctf

CMD socat -dd -v "tcp-listen:7580,reuseaddr,fork" exec:./game.sh
