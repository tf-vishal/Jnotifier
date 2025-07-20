FROM golang:1.24.4-alpine3.22

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o jnotifier ./cmd/bot

CMD ["./jnotifier"]