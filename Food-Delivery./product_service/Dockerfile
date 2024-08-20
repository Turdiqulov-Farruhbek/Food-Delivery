FROM golang:1.22.1 AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o delivery_product ./cmd

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/delivery_product .

COPY .env .env 

RUN chmod +x delivery_product

EXPOSE 6060

CMD ["./delivery_product"]
