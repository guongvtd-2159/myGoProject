FROM golang:1.18-alpine

WORKDIR /app

RUN apk update && apk add --no-cache mysql-client

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o main .

EXPOSE 8080

CMD ["/app/wait-for-mysql.sh", "db", "3306", "/app/main"]

