FROM golang:alpine AS builder

WORKDIR /build

COPY . .
RUN apk add --no-cache git
RUN go mod download

RUN go build -o crm.go_ecommerce_sq ./cmd/server

FROM alpine:latest

COPY ./config /config
COPY --from=builder /build/crm.go_ecommerce_sq /

# Khai báo để ứng dụng có thể đọc biến môi trường trong thời gian chạy
ENV SENDGRID_API_KEY=${SENDGRID_API_KEY}
ENV SENDER_EMAIL=${SENDER_EMAIL}

ENTRYPOINT [ "/crm.go_ecommerce_sq", "config/local.yaml" ]
