# FROM golang:1.22-bookworm AS builder
FROM golang@sha256:96788441ff71144c93fc67577f2ea99fd4474f8e45c084e9445fe3454387de5b AS builder

WORKDIR /workdir

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v .

# FROM debian:bookworm-slim
FROM debian@sha256:67f3931ad8cb1967beec602d8c0506af1e37e8d73c2a0b38b181ec5d8560d395

WORKDIR /workdir

COPY --from=builder /workdir/toybox .

ENTRYPOINT ["./toybox"]
