# Use an official Node.js runtime as a parent image
FROM node:21.6.2-alpine as build-stage

# Creates and Sets the working directory in the container
WORKDIR /app

# Copy package.json and package-lock.json to the working directory
COPY ./client/package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY ./client .

# Build the Vue app for production
RUN npm run build

# Use a smaller image for the production environment
FROM nginx:alpine as production-stage

# Copy the build artifacts from the build-stage to the production-stage
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Expose the port that the app will run on (default is 80 for Vue.js)
EXPOSE 80

# Start Nginx when the container runs
CMD ["nginx", "-g", "daemon off;"]
