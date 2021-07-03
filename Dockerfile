FROM node:alpine AS base

ENTRYPOINT ["tini", "--"]
CMD ["node", "/app/relay.js"]

RUN apk add --no-cache tini

WORKDIR /app
COPY LICENSE package.json relay.js .
RUN npm install
USER node
