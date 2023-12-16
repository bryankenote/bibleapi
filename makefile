BINARY_NAME=bible_server

sqlc:
	rm -rf /src/codegen/sqlc/*
	sqlc generate

pb:
	rm -rf ./src/codegen/pb/*
	buf generate

run:
	go run ./src/main.go

lint:
	buf lint

build:
	go build -o ${BINARY_NAME} ./src/main.go

clean:
	go clean
	rm ${BINARY_NAME}
