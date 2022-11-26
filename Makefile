migrate-create:
	migrate create -ext sql -dir db/migrations -seq init_schema
createdb:
	docker-compose exec -it db createdb --username=root --owner=root simple_bank
migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up