FROM golang:1.22.0-alpine

LABEL authors="Besto"

RUN apk add --no-cache tzdata

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/app/
ENV GIN_MODE=release
ENV FILE_INFO_DB=root:root@tcp(db:3306)/symptom_monitoring_and_treatment_suggestion_file_info?parseTime=true&loc=Asia%2FBangkok
ENV HOST=localhost
ENV PORT=8080
EXPOSE 8080


CMD [ "./main" ]