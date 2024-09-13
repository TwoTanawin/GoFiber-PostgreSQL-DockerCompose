#!/bin/bash

# Base URL for the API
BASE_URL="http://localhost:3000/movies"

# Create a new movie (POST)
echo "Creating a new movie..."
curl -X POST "$BASE_URL" \
     -H "Content-Type: application/json" \
     -d '{"title": "Inception", "year": 2010}' \
     -w "\nResponse Code: %{http_code}\n" \
     -o create_response.json

echo -e "\nCreate response:"
cat create_response.json

# Retrieve all movies (GET)
echo -e "\nRetrieving all movies..."
curl -X GET "$BASE_URL" \
     -w "\nResponse Code: %{http_code}\n" \
     -o all_movies.json

echo -e "\nAll movies response:"
cat all_movies.json

# Extract movie ID from the create response (for subsequent tests)
MOVIE_ID=$(jq -r '.id' create_response.json)
if [ "$MOVIE_ID" == "null" ]; then
  echo "Failed to retrieve movie ID from create response."
  exit 1
fi

# Retrieve a specific movie by ID (GET)
echo -e "\nRetrieving movie with ID $MOVIE_ID..."
curl -X GET "$BASE_URL/$MOVIE_ID" \
     -w "\nResponse Code: %{http_code}\n" \
     -o movie_response.json

echo -e "\nMovie response:"
cat movie_response.json

# Update a specific movie by ID (PUT)
echo -e "\nUpdating movie with ID $MOVIE_ID..."
curl -X PUT "$BASE_URL/$MOVIE_ID" \
     -H "Content-Type: application/json" \
     -d '{"title": "The Dark Knight", "year": 2008}' \
     -w "\nResponse Code: %{http_code}\n" \
     -o update_response.json

echo -e "\nUpdate response:"
cat update_response.json

# Delete a specific movie by ID (DELETE)
echo -e "\nDeleting movie with ID $MOVIE_ID..."
curl -X DELETE "$BASE_URL/$MOVIE_ID" \
     -w "\nResponse Code: %{http_code}\n" \
     -o delete_response.json

echo -e "\nDelete response:"
cat delete_response.json

# Check if the movie is really deleted
echo -e "\nRetrieving all movies after deletion..."
curl -X GET "$BASE_URL" \
     -w "\nResponse Code: %{http_code}\n" \
     -o all_movies_after_delete.json

echo -e "\nAll movies after deletion response:"
cat all_movies_after_delete.json
