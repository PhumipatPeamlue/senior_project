FROM golang:1.22.0-alpine

LABEL authors="Besto"

RUN apk add --no-cache tzdata

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/app/
ENV GIN_MODE=release

ENV REMINDER_DB=root:root@tcp(db:3306)/symptom_monitoring_and_treatment_suggestion_reminders?parseTime=true&loc=Asia%2FBangkok
ENV ADD_NEW_FROM_HOUR_REMINDER_URL=http://notification_service:8080/notification/hour
ENV ADD_NEW_FROM_PERIOD_REMINDER_URL=http://notification_service:8080/notification/period
ENV CHANGE_NOTIFY_TIME_FROM_HOUR_REMINDER_URL=http://notification_service:8080/notification/hour
ENV CHANGE_NOTIFY_TIME_FROM_PERIOD_REMINDER_URL=http://notification_service:8080/notification/period
ENV PORT=8080
EXPOSE 8080


CMD [ "./main" ]