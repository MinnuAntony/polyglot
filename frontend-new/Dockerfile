#
# Stage 1: Build the React application
#
FROM node:20.11-alpine AS build

WORKDIR /app

# Copy package files first for caching
COPY package*.json ./

# Use npm ci for consistent and reproducible builds
RUN npm ci

# Copy all source code, respecting the .dockerignore file
COPY . .

# Build the application for production
RUN npm run build

#
# Stage 2: Serve the application with Nginx
#
FROM nginx:1.25.3-alpine

# Copy custom Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy static files from the build stage
COPY --from=build /app/dist /usr/share/nginx/html

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
