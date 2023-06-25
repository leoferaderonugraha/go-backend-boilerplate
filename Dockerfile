FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .

RUN go mod tidy
RUN go build -ldflags="-s -w" ./cmd/app/main.go

CMD ["./main"]
