FROM ubuntu:20.04

ENV DEBIAN_FRONTEND=noninteractive

RUN apt-get update \
    && apt-get install -y curl gnupg2
RUN curl -fsSL https://deb.nodesource.com/setup_20.x | bash -
RUN apt-get install -y nodejs
RUN apt-get install -y libglib2.0-0 libnss3 libnspr4 libdbus-1-3 libatk1.0-0 libatk-bridge2.0-0 libcups2 libatspi2.0-0 libx11-6 libxcomposite1 libxdamage1 libxext6 libxfixes3 libxrandr2 libgbm1 libdrm2 libxcb1 libxkbcommon0 libpango-1.0-0 libcairo2 libasound2

COPY ./dumb-init_1.2.5_x86_64 /usr/local/bin/dumb-init
RUN chmod +x /usr/local/bin/dumb-init

WORKDIR /app

RUN addgroup appgroup \
    && useradd appuser -G appgroup \
    && mkdir -p /home/appuser/Downloads \
    && chown -R appuser:appgroup /home/appuser \
    && chown -R appuser:appgroup /app

USER appuser

COPY package.json ./

RUN npm i \
    && npx playwright install chromium

COPY . .

ENTRYPOINT ["dumb-init", "--"]
CMD ["node", "index.js"]
