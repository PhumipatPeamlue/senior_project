FROM golang:1.21.5-alpine
LABEL authors="besto"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/document_service/

ENV ES_URL="http://es:9200"
ENV PORT="8080"

EXPOSE 8080

CMD ["./main"]