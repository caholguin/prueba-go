FROM golang:1.21-alpine

RUN apk update && apk upgrade --no-cache

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

CMD ["./main"]
