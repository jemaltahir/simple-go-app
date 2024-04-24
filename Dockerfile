FROM golang:1.22-alpine AS builder

WORKDIR /app

COPY go.mod  ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./...

FROM alpine:latest

COPY --from=builder /app/main /app/main

EXPOSE 8080

WORKDIR /app

CMD ["./main"]

