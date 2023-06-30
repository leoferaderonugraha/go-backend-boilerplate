FROM golang:1.20.5-alpine

WORKDIR /app

COPY . .

RUN go version
RUN go env GOPATH

RUN apk update
RUN apk add --no-cache curl

RUN go mod tidy
RUN go install github.com/cosmtrek/air@latest

# RUN go build -ldflags="-s -w" ./cmd/app/main.go

EXPOSE 3000

CMD ["air", "-c", ".air.toml"]
