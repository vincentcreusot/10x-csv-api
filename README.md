# 10x Genomics Platform Engineering Technical Coding Prompt

## Task 
Create a web service that converts a CSV file into an API that exposes JSON.

We've provided a CSV file of Seattle weather in seattle-weather.csv. It contains the following labels in the header, with the following format:

```
date,precipitation,temp_max,temp_min,wind,weather
...
2012-06-03,0.0,17.2,9.4,2.9,sun
2012-06-04,1.3,12.8,8.9,3.1,rain
...
```

Your tasks (in no specific order):
Read in CSV file, output JSON
Create a server that responds to a GET request with the output in JSON
Set up querying on the data:
Limit results to a number http://my-server.example.com/query?limit=5
By date: http://my-server.example.com/query?date=2012-06-04
By weather type: http://my-server.example.com/query?weather=rain
Create multi-query filtering, eg. http://my-server.example.com/query?weather=rain&limit=5

#### Bonus
Bundle the service into a Docker image and run it as a container
Query the service API from outside the container
#### Bonus 2
Write a script to test the service API
Bundle the test script into its own Docker image
Run the test script image from a separate container and hit the same API

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
docker-compose up
```

