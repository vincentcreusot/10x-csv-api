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