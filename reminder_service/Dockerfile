FROM golang:1.21.5-alpine
LABEL authors="besto"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/reminder_service/

ENV MYSQL_DSN="root:root@tcp(service_db:3306)/service_db?parseTime=true"
ENV PORT="8080"

EXPOSE 8080

CMD ["./main"]