FROM golang:1.22.0-alpine

LABEL authors="Besto"

RUN apk add --no-cache tzdata

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o main ./cmd/app/
ENV REMINDER_DB=root:root@tcp(db:3306)/symptom_monitoring_and_treatment_suggestion_reminders?parseTime=true&loc=Asia%2FBangkok
ENV NOTIFICATION_DB=root:root@tcp(db:3306)/symptom_monitoring_and_treatment_suggestion_notifications?parseTime=true&loc=Asia%2FBangkok
ENV PET_DB=root:root@tcp(db:3306)/symptom_monitoring_and_treatment_suggestion_pets?parseTime=true&loc=Asia%2FBangkok

ENV LINE_ACCESS_TOKEN=gFETqcP4VRi5p/ti4eHBQFL1gnBj3UqWjoyU0u52g3UX20vnmRMjehc3gYyVajPjT8OZ/4sdUxk9kf68bVImo1rMrhj4EwERt8UxWh5+m12c6laBlDkLaNs6y4giHJGn+3VQhf4OFnaTlq48My3VDwdB04t89/1O/w1cDnyilFU=
ENV LINE_MESSAGE_PUSH_URL=https://api.line.me/v2/bot/message/push

CMD [ "./main" ]