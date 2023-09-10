# 10x Genomics Platform Engineering Technical Coding Prompt

## Task 
Create a web service that converts a CSV file into an API that exposes JSON.

## Usage

### Requirements
Uses some tools. A Makefile is provided for the common tasks.

The testing script uses `bash`, `curl` and `jq` for JSON parsing.

### Build
There is a `build` target that builds the go source code to a `10x-csv-api` binary.
That means you can run: 
```
make build
```
to build the go binary.

### Run
The executable takes the CSV file path as command line argument. It will start a server on port 8080.
A target exists that takes the provided CSV file and runs the executable. You can then execute: 
```
make run
```


### Docker build
A multistage `dockerfile` is provided that builds the go package in a container and run the command in another to reduce the image size. Building the image is done with
```
make docker-build
```
### Docker run
The image can be ran using
```
make docker-run
```

## Tests
### Unit tests
The csv parser contains unit tests. You can run them using:
```
make test
```
### API tests
Tests for the API are done with a bash script. The script uses an argument to specify the host.

If you run the docker container with the test script, you can pass `localhost` as argument.
```
bash test/test.sh localhost
```
or with: 
```
make api-test
```

### Docker
The `Dockerfile.test` file provided uses the `test.sh` bash script.

The docker image can be run with the API container with a `docker-compose` definition.
The composed package can be run using 
```
make compose
```
Then to ensure the containers are not running, run 
```
make compose-down
```


