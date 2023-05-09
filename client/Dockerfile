# Build the react client
FROM node:18-alpine3.17 AS build
WORKDIR /app
COPY . ./
RUN npm install
RUN npm run build

# Start NGINX
FROM nginx:stable-alpine3.17
COPY --from=build /app/dist /usr/share/nginx/html
COPY ./nginx.conf /etc/nginx/nginx.conf
EXPOSE 80
CMD [ "nginx", "-g", "daemon off;" ]
