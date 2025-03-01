FROM golang:1.24.0-alpine AS builder
RUN apk add --no-cache git ca-certificates && update-ca-certificates

COPY . .
RUN go mod download && \
    GOARCH=amd64 CGO_ENABLED=0 GOOS=linux go build -o /app ./main.go

FROM scratch AS runner
COPY --from=builder /app /app
EXPOSE 8080

CMD ["/app"]