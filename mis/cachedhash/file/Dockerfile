# FROM golang:1.22-alpine AS builder
FROM golang@sha256:6522f0ca555a7b14c46a2c9f50b86604a234cdc72452bf6a268cae6461d9000b AS builder

WORKDIR /usr/src/cachedhash

COPY go.mod go.sum ./
RUN go mod download && go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/cachedhash ./...
RUN cachedhash -mode hash -file ./flag.txt > ./hashed_flag.txt

# FROM gcr.io/distroless/static-debian12:nonroot
FROM gcr.io/distroless/static-debian12@sha256:e9ac71e2b8e279a8372741b7a0293afda17650d926900233ec3a7b2b7c22a246

COPY --from=builder /usr/local/bin/cachedhash /usr/local/bin/
COPY --from=builder /usr/src/cachedhash/hashed_flag.txt /

ENTRYPOINT ["cachedhash", "-mode", "serve", "-file", "/hashed_flag.txt"]
