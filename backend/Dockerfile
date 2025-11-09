FROM golang:1.25.3-alpine3.22 AS builder
WORKDIR /usr/src/app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o application ./cmd/app

FROM alpine:3.22
WORKDIR /app

COPY --from=builder /usr/src/app/config.yaml .
COPY --from=builder /usr/src/app/application ./app
COPY migrations migrations

EXPOSE 8080

CMD ["./app"]
