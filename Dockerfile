FROM golang:1.23.0-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o habit-tracker .

EXPOSE 8080

CMD ["./habit-tracker"]

