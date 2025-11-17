FROM golang:1.25.3 AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN GOOS=linux go build -o todo_server .

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y ca-certificates && rm -rf /var/lib/apt/lists/*
WORKDIR /app
COPY --from=builder /app/todo_server .
COPY --from=builder /app/web ./web
ENV TODO_PORT=7540
ENV TODO_DBFILE=scheduler.db
ENV TODO_PASSWORD="12345"

EXPOSE 7540
CMD ["./todo_server"]
