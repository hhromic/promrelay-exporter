FROM node:16.4-alpine

LABEL org.opencontainers.image.source https://github.com/hhromic/prometheus-relay-exporter

ENV NODE_ENV=production
ENTRYPOINT ["tini", "--"]
CMD ["node", "server.js"]
EXPOSE 8080

RUN apk add --no-cache tini

WORKDIR /app
COPY package*.json ./
RUN npm ci --only=production

COPY server.js ./
USER node
