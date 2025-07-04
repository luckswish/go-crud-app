FROM golang:1.22-alpine

RUN apk add --no-cache git

WORKDIR /app

# 👇 Сначала копируем только go.mod и go.sum
COPY go.mod go.sum ./

# ✅ Качаем зависимости
RUN go mod download

# 👇 Потом копируем весь остальной проект
COPY . .

# 👇 Собираем
RUN go build -o main ./cmd

CMD ["./main"]
