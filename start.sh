#!/bin/sh

set -e 

echo "run db migration"
/app/migrate -path /app/migration -database postgresql://root:mokopass@localhost:5432/go_native?sslmode=disable -verbose up

echo "start the app"
exec "$@"