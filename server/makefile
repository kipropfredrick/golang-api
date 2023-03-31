postgresinit:
	docker run --name postgres15 -p 5433:5433 -e POSTGRES_USER=godbtest -e POSTGRES_PASSWORD=31877101 -d postgres:15-alpine

postgres:
	docker exec -it postgres15 psql

createdb:
	docker exec -it postgres15 createdb --username=godbtest --owner=godbtest postgres

dropdb:
	docker exec -it postgres15 dropdb postgres

migrateup:
	migrate -path db/migrations -database "postgresql://godbtest:31877101@go-chat.cydlpjqntwzu.us-west-2.rds.amazonaws.com:5432/postgres?sslmode=disable"  up

migratedown:
	migrate -path db/migrations -database "postgresql://godbtest:31877101@go-chat.cydlpjqntwzu.us-west-2.rds.amazonaws.com:5432/postgres?sslmode=disable" -verbose down

.PHONY: postgresinit postgres createdb dropdb migrateup migratedown