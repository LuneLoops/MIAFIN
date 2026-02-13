# Build stage
FROM oven/bun:1 AS builder

WORKDIR /app

# Copiar package files
COPY package.json bun.lock* ./
RUN bun install

# Copiar código fuente
COPY . .

# Build
RUN bun run build

# Run stage
FROM nginx:alpine

# Copiar archivos de build
COPY --from=builder /app/dist /usr/share/nginx/html

# Configuración de nginx
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
