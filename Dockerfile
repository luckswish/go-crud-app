FROM golang:1.23-alpine

RUN apk add --no-cache git

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

# Собираем два отдельных бинарника
RUN go build -o main ./cmd
RUN go build -o migrate ./cmd/migrate

CMD ["./main"]
