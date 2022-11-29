include .env

build:
	go build -o .bin/main cmd/main/app.go

prod: build
	./.bin/main production

dev:
	go run cmd/main/app.go

db_up:
	migrate -path ./schema -database "postgresql://postgres:po_psql@localhost:5432/wallet?sslmode=disable" up

run_sub:
	go run cmd/sub/app.go

db_down:
	migrate -path ./schema -database "postgresql://postgres:po_psql@localhost:5432/wallet?sslmode=disable" down