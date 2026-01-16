# build environment
FROM node:slim as build
WORKDIR /app
COPY . .
RUN yarn
RUN yarn build:dev

# production environment
FROM nginx:stable
COPY --from=build /app/dist /usr/share/nginx/html
COPY --from=build /app/nginx.conf /etc/nginx/conf.d/default.conf
EXPOSE 80
CMD ["nginx", "-g", "daemon off;"]
