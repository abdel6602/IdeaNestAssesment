FROM golang:1.20

WORKDIR /app

COPY ../go.mod .
COPY ../pkg/api/routes/routes.go .

RUN go run routes.go
