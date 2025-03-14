FROM golang:1.24.1-alpine

WORKDIR /app

RUN apk add --no-cache poppler-utils

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./main.go

EXPOSE 8080

CMD ["./main"]
