FROM golang:1.22.0-alpine

LABEL authors="Besto"

RUN apk add --no-cache tzdata

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/app/
ENV GIN_MODE=release
ENV DB_DSN=root:root@tcp(db:3306)/symptom_monitoring_and_treatment_suggestion_user_web?parseTime=true&loc=Asia%2FBangkok
ENV PORT=8080
ENV TZ=Asia/Bangkok
EXPOSE 8080


CMD [ "./main" ]