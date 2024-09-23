FROM golang:1.22.2 as builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 go build -o welcom main.go

FROM alpine:3.15.3
WORKDIR /app
COPY --from=builder /app/welcome .
COPY --from=builder /app/static .

CMD ["./welcome"]