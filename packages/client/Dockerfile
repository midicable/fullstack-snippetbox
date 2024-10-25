FROM node:current-alpine AS builder
WORKDIR /app
COPY . .
RUN yarn
RUN yarn build

FROM nginx:1.27.1-alpine
COPY nginx.conf /etc/nginx/conf.d/default.conf
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 3000
