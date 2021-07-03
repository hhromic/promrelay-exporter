FROM node:alpine AS base

LABEL org.opencontainers.image.source https://github.com/hhromic/prometheus-relay

ENTRYPOINT ["tini", "--"]
CMD ["node", "/app/relay.js"]

RUN apk add --no-cache tini

WORKDIR /app
COPY LICENSE package.json relay.js .
RUN npm install
USER node
