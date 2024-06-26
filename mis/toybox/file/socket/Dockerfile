# FROM buildpack-deps:bookworm AS builder
FROM buildpack-deps@sha256:d56cd472000631b8faca51f40d4e3f1b20deffa588f9f207fa6c60efb62ba7c4 AS builder

RUN apt-get update && \
		apt-get install -y cmake libseccomp-dev && \
		rm -rf /var/lib/apt/lists/*

WORKDIR /workdir

COPY . .

RUN cmake -DCMAKE_BUILD_TYPE=Release -B cmake-build && \
    cmake --build cmake-build --target all

# FROM debian:bookworm-slim
FROM debian@sha256:67f3931ad8cb1967beec602d8c0506af1e37e8d73c2a0b38b181ec5d8560d395

RUN apt-get update && \
		apt-get install -y socat && \
		rm -rf /var/lib/apt/lists/*

WORKDIR /workdir

COPY banner.txt flag.txt ./
COPY --from=builder /workdir/cmake-build/server /workdir/cmake-build/sandbox ./

USER nobody

ENTRYPOINT ["socat", "-dd", "-v", "tcp-listen:9893,reuseaddr,fork", "exec:./server,stderr"]
