# Build stage
FROM node:20-alpine AS build
WORKDIR /app
RUN corepack enable

ARG VITE_PUBLIC_SITE_URL="https://one-earth.info"
ENV VITE_PUBLIC_SITE_URL=$VITE_PUBLIC_SITE_URL

COPY package.json .
COPY .yarnrc.yml .
COPY yarn.lock* .
RUN if [ -f yarn.lock ]; then yarn install --immutable; else yarn install; fi

COPY . .
RUN yarn build

FROM nginx:alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=build /app/dist /usr/share/nginx/html
