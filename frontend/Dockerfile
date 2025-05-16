# Development stage (for hot-reloading)
FROM node:18 as dev
WORKDIR /app
# First copy only package files for better caching
COPY package*.json ./
RUN npm install

# Then copy everything else
COPY . .

# Configure Vite for Docker compatibility
ENV VITE_API_URL=http://backend:8080
EXPOSE 5173

# Run with polling for file changes in Docker
CMD ["npm", "run", "dev", "--", "--host", "0.0.0.0"]

# Builder stage (for production)
FROM node:18 as builder
WORKDIR /app
COPY --from=dev /app .
RUN npm run build

# Production stage
FROM nginx:alpine as prod

# Remove default nginx configs
RUN rm -rf /etc/nginx/conf.d/default.conf
COPY nginx.conf /etc/nginx/conf.d

# Copy built files
COPY --from=builder /app/dist /usr/share/nginx/html

EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]