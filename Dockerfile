FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o main cmd/main.go

# Run stage
# FROM gcr.io/distroless/base-debian12
FROM alpine:3.19
WORKDIR /app
COPY --from=builder /app/main .
COPY ./config.env .
# COPY start.sh .
# COPY wait-for.sh .
COPY db/migration db/migration

EXPOSE 8080 9090
CMD [ "/app/main" ]
# ENTRYPOINT [ "/app/start.sh" ]