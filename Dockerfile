FROM golang:alpine AS builder

WORKDIR /build

COPY . .

RUN go mod download

RUN go build -o crm.go_ecommerce_sq ./cmd/server

FROM scratch

COPY ./config /config

COPY --from=builder /build/crm.go_ecommerce_sq /

ENTRYPOINT [ "/crm.go_ecommerce_sq", "config/local.yaml" ]