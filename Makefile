BINARY_NAME=10x-csv-api
all: build test

.PHONY: build test run docker-run docker-build
build: deps
	@go build -o ${BINARY_NAME} cmd/10x-csv-api/10x-csv-api.go

deps:
	@go mod download

test:
	@go test -v ./...

api-test: 
	bash test/test.sh localhost
 
run: build
	./${BINARY_NAME} seattle-weather.csv

docker-build:
	@docker build -t ${BINARY_NAME} -f Dockerfile .

docker-run: docker-build
	@docker run -p 8080:8080 ${BINARY_NAME}

compose-build:
	@docker-compose build

compose: compose-build
	@docker-compose up

compose-down:
	@docker-compose down

clean:
	@go clean
	@rm -f ${BINARY_NAME}
	@docker rmi ${BINARY_NAME}
