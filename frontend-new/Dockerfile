# Stage 1: Build the React application
FROM node:20-alpine AS build
WORKDIR /app

# Copy package.json and package-lock.json to leverage Docker layer caching
COPY polyglot-frontend/package*.json ./

# Install project dependencies
RUN npm install

# Copy all the rest of the source code
COPY polyglot-frontend/ .

# Build the app for production. This creates a 'dist' folder with static files.
RUN npm run build

# Stage 2: Serve the application with Nginx
FROM nginx:alpine

# Copy the custom Nginx configuration file
COPY ./nginx.conf /etc/nginx/conf.d/default.conf

# Copy the static build output from the 'build' stage into the Nginx public directory
COPY --from=build /app/dist /usr/share/nginx/html

# Expose port 80 to the host machine
EXPOSE 80

# Start Nginx when the container launches
CMD ["nginx", "-g", "daemon off;"]
