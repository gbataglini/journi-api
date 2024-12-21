FROM golang:1.23 AS builder

COPY . .

RUN CGO_ENABLED=0 go build -o /app ./cmd/journi-api/main.go

FROM alpine 

COPY --from=builder /app app

CMD ["./app"]