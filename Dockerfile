FROM golang:1.23.1 AS builder

WORKDIR /app

COPY . .

RUN go mod download
RUN go build -o proxy-server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/proxy-server .

EXPOSE 8080

CMD ["./proxy-server"]
