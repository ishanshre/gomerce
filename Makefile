#!make

include .env

DB_URL=postgresql://${db_username}:${db_password}@localhost:5432/${db_dbname}?sslmode=disable

run:
	go run ./cmd/api

help:
	go run ./cmd/api -h

createDBContainer:
	docker run --name gomerce -e POSTGRES_USER=${db_username} -e POSTGRES_PASSWORD=${db_password} -p 5432:5432 -d postgres

createDBPGadmin4Container:
	docker run --name gomercePgadmin -p 5050:80 -e 'PGADMIN_DEFAULT_EMAIL=admin@admin.com' -e 'PGADMIN_DEFAULT_PASSWORD=admin' -d dpage/pgadmin4

createRedisContainer:
	docker run -d --name gomerceRedis -p 6379:6379 redis:latest 

startContainer:
	docker start gomerce gomercePgadmin gomerceRedis

stopContainer:
	docker stop gomerce gomercePgadmin gomerceRedis
	
migrateUp: 
	migrate -path migrations -database "${DB_URL}" -verbose up

migrateDown: 
	migrate -path migrations -database "${DB_URL}" -verbose down

migrateForce: 
	migrate -path migrations -database "${DB_URL}" force $(version)

migrateCreate:
	migrate create -ext sql -dir migrations -seq $(fileName)
