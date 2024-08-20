FROM golang:1.22.1 AS builder

WORKDIR /app

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o delivery_gateway .

FROM alpine:latest

RUN apk add --no-cache ca-certificates

WORKDIR /app

COPY --from=builder /app/delivery_gateway .
COPY --from=builder /app/api/casbin/model.conf ./api/casbin/
COPY --from=builder /app/api/casbin/policy.csv ./api/casbin/
COPY .env .env 

RUN chmod +x delivery_gateway

EXPOSE 8090

CMD ["./delivery_gateway"]