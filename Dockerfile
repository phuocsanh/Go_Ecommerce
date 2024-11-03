# FROM golang:alpine AS builder

# WORKDIR /build

# COPY . .
# RUN apk add --no-cache git
# RUN go mod download

# RUN go build -o crm.go_ecommerce_sq ./cmd/server

# FROM alpine:latest

# COPY ./config /config
# COPY --from=builder /build/crm.go_ecommerce_sq /

# # Khai báo để ứng dụng có thể đọc biến môi trường trong thời gian chạy
# ENV SENDGRID_API_KEY=${SENDGRID_API_KEY}
# ENV SENDER_EMAIL=${SENDER_EMAIL}

# ENTRYPOINT [ "/crm.go_ecommerce_sq", "config/local.yaml" ]

FROM golang:alpine

# Cài đặt các công cụ cần thiết
RUN apk add --no-cache git

# Thiết lập thư mục làm việc
WORKDIR /app

# Copy toàn bộ mã nguồn vào container
COPY . .

# Tải xuống các dependencies của Go
RUN go mod download

# Expose port (nếu ứng dụng của bạn chạy trên một port cụ thể)
EXPOSE 8002

# Sử dụng lệnh go run để phát triển mà không cần build lại toàn bộ
CMD ["go", "run", "./cmd/server"]
