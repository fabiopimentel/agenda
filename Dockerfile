FROM golang:1.14.2-alpine3.11
RUN mkdir app
ADD . app/
WORKDIR app
RUN go build -o main ./cmd/server/main.go
EXPOSE 5000
CMD ["/app/main"]
