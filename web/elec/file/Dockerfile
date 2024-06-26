# FROM golang:1.22.3-bookworm AS server-builder
FROM golang@sha256:5c56bd47228dd572d8a82971cf1f946cd8bb1862a8ec6dc9f3d387cc94136976 AS server-builder

WORKDIR /root/elec-server

COPY elec-server/go.mod elec-server/go.sum ./
RUN go mod download && go mod verify

COPY elec-server .
RUN CGO_ENABLED=1 go build -v .

# FROM node:20.13.1-bookworm AS admin-console-builder
FROM node@sha256:d6925dc84f8c0d1c1f8df4ea6a9a54e57d430241cb734b1b0c45ed6d26e8e9c0 AS admin-console-builder

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update -y && \
    apt-get install -y --no-install-recommends dpkg fakeroot && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /root/elec-admin-console

COPY elec-admin-console/package.json elec-admin-console/package-lock.json  ./
RUN npm ci

COPY elec-admin-console .
RUN npm run make

# FROM buildpack-deps:bookworm-curl
FROM buildpack-deps@sha256:26f797741018701856083d5936027cf8433ec952fac9be25a45097c1ed5661d3

RUN export DEBIAN_FRONTEND=noninteractive && \
    apt-get update -y && \
    apt-get install -y --no-install-recommends xvfb xauth libgtk-3-0 libnotify4 libnss3 libxtst6 xdg-utils libatspi2.0-0 libdrm2 libgbm1 libxcb-dri3-0 libglib2.0-bin libasound2 && \
    rm -rf /var/lib/apt/lists/*

COPY --from=admin-console-builder /root/elec-admin-console/out/make/deb/x64/elec-admin-console_0.0.1_amd64.deb /root/elec-admin-console.deb
RUN dpkg -i /root/elec-admin-console.deb && \
    rm /root/elec-admin-console.deb

COPY --from=server-builder /root/elec-server/elec-server /usr/local/bin/elec-server

COPY flag /flag

RUN useradd --create-home elec
USER elec
WORKDIR /home/elec
RUN mkdir data

CMD rm -f /tmp/.X99-lock && xvfb-run /usr/local/bin/elec-server
