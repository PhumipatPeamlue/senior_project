FROM golang:1.22.0-alpine

LABEL authors="Besto"

RUN apk add --no-cache tzdata

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/app/
ENV GIN_MODE=release
ENV ES_URL=http://es:9200
ENV VIDEO_DOC_INDEX=video_doc
ENV DRUG_DOC_INDEX=drug_doc
ENV PORT=8080
EXPOSE 8080


CMD [ "./main" ]