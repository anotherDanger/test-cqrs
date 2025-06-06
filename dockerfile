FROM golang:1.24.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

COPY src /app/src

RUN go build -o main ./src/App

FROM alpine:latest

COPY --from=builder /app/main .
COPY --from=builder /app/src/App/logs /app/src/App/logs

CMD [ "./main" ]