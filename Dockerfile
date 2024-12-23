FROM golang:1.22 AS builder
WORKDIR /usr/src/app
ENV GOPROXY=https://goproxy.io,direct
COPY . .
RUN CGO_ENABLED=1 GOOS=linux go build -a -installsuffix cgo -o order ./cmd/main.go

FROM debian:latest
COPY --from=builder /usr/src/app/order ./order
CMD ["./order"]