BINARY_NAME=bible_server

sqlc:
	rm -rf ./codegen/sqlc/*
	sqlc generate

pb:
	rm -rf ./codegen/pb/*
	buf generate

run:
	go run ./cmd/main.go

lint:
	buf lint

build:
	go build -o ${BINARY_NAME} ./cmd/main.go

clean:
	go clean
	rm ${BINARY_NAME}
