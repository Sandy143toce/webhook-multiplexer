# Dockerfile
FROM golang:1.21.4-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build -o main .

EXPOSE 8000

CMD ["./main"]
