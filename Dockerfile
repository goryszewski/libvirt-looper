FROM golang:1.22.0 as builder

WORKDIR /app

RUN apt update && apt  install libvirt-clients libvirt-dev -y

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o looper .

FROM alpine:3.6

WORKDIR /app

COPY --from=builder /app/looper /app/looper

CMD ["./looper"]