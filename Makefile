createpg:
	docker run --name some-postgres --network blognet -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=password -d postgres:14-bookworm

runpg:
	docker start some-postgres && docker ps

intopg:
	docker exec -it some-postgres psql -U root


stoppg:
	docker stop some-postgres

createdb:
	docker exec -it some-postgres createdb --username=root --owner=root blog

dropdb:
	docker exec -it some-postgres dropdb blog 

createmigration:
	migrate create -ext sql -dir db/migration -seq $(name) -verbose

migrateup:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/blog?sslmode=disable" -verbose up $(step)


migratedown:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/blog?sslmode=disable" -verbose down $(step)


migrateforce:
	migrate -path db/migration -database "postgresql://root:password@localhost:5432/blog?sslmode=disable" -verbose force $(version)

sqlc:
	sqlc generate

test:
	go test  -count=1 -timeout 120s -p 1 ./...

server:
	go run main.go

mock:
	mockgen -package mockdb -destination db/mock/store.go github.com/Shubham-Rasal/blog-backend/db/sqlc Store

.PHONY: createdb dropdb createpg runpg  stoppg migrateup migratedown test sqlc server mock intopg migrateforce createmigration 