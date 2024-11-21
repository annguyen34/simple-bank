# Build stage
FROM golang:1.21.0-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz  | tar xvz

# Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/migrate ./migrate
COPY app.env .
COPY start.sh .
RUN chmod +x start.sh
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8000
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]