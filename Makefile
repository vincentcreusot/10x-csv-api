BINARY_NAME=10x-csv-api
all: build test

.PHONY: build test run docker-run docker-build
build:
	go build -o ${BINARY_NAME} cmd/10x-csv-api/10x-csv-api.go
 
test:
	go test -v ./...
 
run: build
	./${BINARY_NAME} seattle-weather.csv

docker-build:
    docker build -t ${BINARY_NAME} .

docker-run:
    docker run -p 8080:8080 ${BINARY_NAME} seattle-weather.csv
 
clean:
	go clean
	rm -f ${BINARY_NAME}
	docker rmi ${BINARY_NAME}
