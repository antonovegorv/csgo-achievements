build:
	go build -o ./bin/matchestracker ./cmd/matchestracker/main.go
	go build -o ./bin/gcclient ./cmd/gcclient/main.go
	go build -o ./bin/demodownloader ./cmd/demodownloader/main.go

test:
	go test -v -race -timeout 30s ./...

migrate_up:
	migrate -path migrations -database "postgres://postgres:postgrespw@localhost:55000/postgres?sslmode=disable" up

migrate_down:
	migrate -path migrations -database "postgres://postgres:postgrespw@localhost:55000/postgres?sslmode=disable" down

migrate_test_up:
	migrate -path migrations -database "postgres://postgres:postgrespw@localhost:55000/postgres_test?sslmode=disable" up

migrate_test_down:
	migrate -path migrations -database "postgres://postgres:postgrespw@localhost:55000/postgres_test?sslmode=disable" down
