#!/bin/bash
host=$1
API_URL="http://$host:8080/query"

function test_api() {

  echo "Testing API..."

  # Test API response with limit 1
  response=$(curl -s "$API_URL?limit=1")
  echo "Response: $response"

  # Test no filters
  response=$(curl -s "$API_URL")
  check_response_length $response 1461

  # Test date filter
  response=$(curl -s "$API_URL?date=2012-01-01")
  check_response_length $response 1

  # Test weather filter 
  response=$(curl -s "$API_URL?weather=rain")
  check_response_length $response 641

  # Test limit filter
  response=$(curl -s "$API_URL?limit=2")
  check_response_length $response 2

  # Test multiple filters
  response=$(curl -s "$API_URL?date=2012-01-01&weather=rain&limit=1")
  check_response_length $response 0

  echo "API tests completed"
}

function check_response_length() {
  length=$(echo $1 | jq length)
  expected=$2

  if [ "$length" != "$expected" ]; then
    echo "Response length $length does not match expected $expected"
    exit 1
  else
    echo "Response length matches expected"
  fi
}

test_api
