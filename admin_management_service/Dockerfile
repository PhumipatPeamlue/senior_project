FROM golang:1.21.5-bookworm
LABEL authors="besto"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o main ./cmd/rest/

ENV ES_URL="http://es:9200"
ENV MYSQL_DSN="root:root@tcp(image_info_db:3306)/image_info_db?parseTime=true"
ENV VIDEO_DOC_INDEX="video_doc"
ENV DRUG_DOC_INDEX="drug_doc"
ENV HOST="localhost"
ENV PORT="8080"
ENV STORAGE_PATH="uploads"

EXPOSE 8080

CMD ["./main"]