FROM golang:1.14.2-alpine3.11 as builder
WORKDIR /app
COPY . /app
RUN go build -o main ./cmd/server/main.go

FROM alpine:latest
WORKDIR app/
COPY --from=builder /app/main .
EXPOSE 5000
CMD ["./main"]

