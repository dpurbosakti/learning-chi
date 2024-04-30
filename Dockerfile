FROM golang:1.21 AS builder

WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Run stage
# WORKDIR /app
# COPY --from=builder /app/main .
COPY cmd/config.env .
COPY start.sh .
COPY wait-for.sh .
COPY migration ./migration

EXPOSE 8080 9090
CMD [ "/app/main" ]
ENTRYPOINT [ "/app/start.sh" ]