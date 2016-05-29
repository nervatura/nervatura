FROM alpine
EXPOSE 3000
EXPOSE 5858

RUN apk update && apk upgrade
RUN apk add --no-cache \
  nodejs \
  python \
  git

COPY . /app/dist
RUN mkdir app/src
RUN mkdir app/data
RUN mkdir app/data/database
RUN cp /app/dist/data/database/demo.db /app/data/database/demo.db 

RUN cd /app/dist; \
  npm install --production --save; \
  npm install nodemailer --save; \
  mv node_modules ../node_modules;
RUN npm install -g pm2@latest

ENV HOST_TYPE=docker