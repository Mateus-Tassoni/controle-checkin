FROM golang:alpine AS builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o api-checkin ./cmd/api/main.go

FROM alpine:latest

RUN apk add --no-cache tzdata
WORKDIR /root/
COPY --from=builder /app/api-checkin .
COPY --from=builder /app/.env . 
EXPOSE 8080
CMD ["./api-checkin"]