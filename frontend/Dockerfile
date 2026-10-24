# Stage 1: Build static assets
FROM node:24-alpine AS builder
WORKDIR /app

# Copy package files and install dependencies
COPY package*.json ./
RUN npm ci

# Copy the rest of your frontend code and build
COPY . .
RUN npm run build

# Stage 2: Serve files with Nginx
FROM nginx:1.30.6-alpine3.24

# 1. Copy your custom Nginx configuration into the container
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy the built vanilla JS files from Vite's default output directory (dist)
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 8888
CMD ["nginx", "-g", "daemon off;"]
