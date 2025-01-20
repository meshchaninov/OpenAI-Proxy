FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY . .
RUN go build -o proxy-server

FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/proxy-server .

EXPOSE 8080

CMD ["./proxy-server"]
