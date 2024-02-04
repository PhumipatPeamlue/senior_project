FROM node:21
LABEL authors="besto"

WORKDIR /usr/src/app

COPY package*.json ./

RUN npm install

COPY . .

ENV DATABASE_HOST="service_db"
ENV DATABASE_NAME="service_db"

CMD ["node", "index.js"]