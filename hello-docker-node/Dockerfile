FROM node:21-alpine3.19 as builder

WORKDIR /usr/src/app

COPY src/ /usr/src/app
RUN npm ci
RUN npm run build

FROM node:21-alpine3.19
WORKDIR /app


ENV DEFAULT_NAME=Manuel
EXPOSE 3000

COPY --chown=nobody:nobody --from=builder /usr/src/app/package.json /usr/src/app/package-lock.json /app/
RUN npm i --only=production

COPY --chown=nobody:nobody --from=builder /usr/src/app/index.js /app

USER nobody

CMD ["node", "/app/index.js"]
