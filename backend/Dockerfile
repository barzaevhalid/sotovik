# ----------- Build stage ------------
FROM golang:1.22 AS builder

WORKDIR /app

# копируем go.mod и go.sum
COPY go.mod go.sum ./

# скачиваем зависимости (кэшируется)
RUN go mod download

# копируем весь проект
COPY . .

# собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o server ./cmd/api

# ----------- Run stage ------------
FROM alpine:latest

WORKDIR /root/

# добавляем сертификаты (чтобы работал HTTPS)
RUN apk add --no-cache ca-certificates

# копируем бинарник из первого этапа
COPY --from=builder /app/server .

# порт твоего Fiber (например 8080)
EXPOSE 8080

# запускаем приложение
CMD ["./server"]
