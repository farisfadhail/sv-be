FROM golang:1.24.3-alpine AS builder

RUN apk add --no-cache git gcc musl-dev

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/google/wire/cmd/wire@latest

WORKDIR /app/internal/injector
RUN wire

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

RUN if [ -f "cmd/migration.go" ]; then \
    CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o migrate cmd/migration.go; \
    fi

FROM alpine:latest

RUN apk --no-cache add ca-certificates tzdata

ENV TZ=Asia/Jakarta

WORKDIR /root/

COPY --from=builder /app/main .

COPY --from=builder /app/migrate* ./

COPY --from=builder /app/.env* ./

COPY --from=builder /app/database/migrations ./database/migrations/

COPY migrate.sh .
RUN chmod +x migrate.sh

EXPOSE 8080

ENTRYPOINT ["./entrypoint.sh"]