compile:
	env GOOS=linux go build -ldflags="-s -w" -o bin/transactionsroutines cmd/http/*.go
	chmod 0777 bin/* -v
run:
	make compile
	./bin/transactionsroutines
clearconfigs:
	rm -rf ./bin -v
test:
	go test -coverpkg=./... ./...
test-report:
	go test -coverpkg=./... -coverprofile=coverage.out -covermode=count ./...
	go tool cover -html=coverage.out
run-docker:
	docker-compose up -d
run-docker-clean:
	docker-compose build --no-cache
	make run-docker
migrate:
	docker exec -i transactions-routines_db_1 mysql -uroot -proot < ./scripts/database.sql