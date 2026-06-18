FROM node:20-slim AS base
WORKDIR /app
COPY package*.json ./
COPY tsconfig.json ./
COPY next.config.ts ./

FROM base AS dev
RUN npm install
COPY . .
EXPOSE 3000
CMD npm run dev

FROM base AS builder
RUN npm ci
COPY . .
RUN npm run build

FROM node:20-alpine AS prod
WORKDIR /app

COPY package*.json ./
RUN npm ci --omit=dev

COPY --from=builder /app/.next ./.next
COPY --from=builder /app/public ./public
COPY --from=builder /app/next.config.ts ./next.config.ts

USER node
EXPOSE 3000
CMD npm run start
