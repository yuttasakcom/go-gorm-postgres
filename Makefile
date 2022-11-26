migrate-create-db:
	docker-compose exec -it db createdb --username=root --owner=root simple_bank

migrate-create-schema:
	migrate create -ext sql -dir db/migrations -seq init_schema

migrate-create-seeder:
	migrate create -ext sql -dir db/migrations -seq seeder

migrateup:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose up
	
migratedown:
	migrate -path db/migrations -database "postgresql://root:secret@localhost:5432/simple_bank?sslmode=disable" -verbose down

test:
	go test -v -cover ./...