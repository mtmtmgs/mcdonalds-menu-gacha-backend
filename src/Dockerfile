# ビルドコンテナ
FROM golang:1.21.4-alpine3.18 as builder

RUN mkdir -p /src

WORKDIR /src

COPY . .

RUN go mod tidy

RUN mkdir -p /app && \
    go build -ldflags="-w -s" -o /app/built-app

# ランタイムコンテナ
FROM alpine

WORKDIR /app

COPY --from=builder /app /app

EXPOSE 8080

CMD ["./built-app"]
