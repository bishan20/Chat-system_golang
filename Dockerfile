#build stage
FROM golang:1.20.2-alpine3.16 AS builder
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN go install github.com/pressly/goose/v3/cmd/goose@latest
RUN which goose

#Run stage
FROM alpine:3.16
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /go/bin/goose ./migrate
COPY app.env .
COPY start.sh .
COPY wait-for.sh .
COPY db/migration ./migration

EXPOSE 8080
CMD [ "/app/main" ]
RUN chmod +x /app/start.sh
ENTRYPOINT [ "/app/start.sh" ]