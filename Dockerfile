FROM golang:1.21-alpine3.18 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY ./ ./
RUN CGO_ENABLED=0 GOOS=linux go build -o htmxtodo

FROM alpine:3.18

WORKDIR /app

ENV MODE=prod
ENV BE_HOST=0.0.0.0
ENV BE_PORT=80

COPY --from=builder /app/htmxtodo /app/public /app/views ./

EXPOSE 80
CMD ./htmxtodo
