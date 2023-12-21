FROM golang:1.21.5-bookworm

RUN apt-get update && apt-get install -y netcat-openbsd

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN chmod +x ./scripts/wait_es.sh

RUN go build -o main ./cmd/main

ENV ES_URL=http://elasticsearch:9200
ENV MYSQL_DSN=root:root@tcp(mysql:3306)/senior_project

EXPOSE 8080

CMD ["./scripts/wait_es.sh", "elasticsearch", "9200", "./main"]