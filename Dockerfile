FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o dungeon ./cmd

FROM alpine:latest
WORKDIR /app
COPY --from=builder ./app/dungeon .
COPY data ./data
CMD ["./dungeon"]