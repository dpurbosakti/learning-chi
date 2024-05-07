DB_URL=postgresql://root:mokopass@localhost:5432/go_native?sslmode=disable

## run the server
run:
	go run cmd/main.go

## postgresrun: start docker postgres
postgresrun:
	docker start postgres14

## createdb: create database
createdb:
	docker exec -it postgres14 createdb --username=root --owner=root go_native

## dropdb: drop database
dropdb:
	docker exec -it postgres14 dropdb go_native

## migrateup: migrate all up schema sql
migrateup:
	migrate -path db/migration -database "$(DB_URL)" -verbose up

## migrateup1: migrate up schema sql by 1
migrateup1:
	migrate -path db/migration -database "$(DB_URL)" -verbose up 1

## migratedown: migrate all down schema sql
migratedown:
	migrate -path db/migration -database "$(DB_URL)" -verbose down

## migratedown1: migrate down schema sql by 1
migratedown1:
	migrate -path db/migration -database "$(DB_URL)" -verbose down 1

## new_migration: init sql migration file with name as parameter
new_migration:
	migrate create -ext sql -dir db/migration -seq $(name)

## sqlc: generate repository code from query
sqlc:
	sqlc generate