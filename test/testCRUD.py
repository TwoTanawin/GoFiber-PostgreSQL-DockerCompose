import requests

BASE_URL = "http://localhost:3000"

def create_movie(title, year):
    url = f"{BASE_URL}/movies"
    payload = {
        "title": title,
        "year": year
    }
    response = requests.post(url, json=payload)
    if response.status_code == 201:
        print("Movie created successfully:")
        print(response.json())
    else:
        print("Failed to create movie:")
        print(response.status_code, response.text)

def get_movies():
    url = f"{BASE_URL}/movies"
    response = requests.get(url)
    if response.status_code == 200:
        print("Movies retrieved successfully:")
        print(response.json())
    else:
        print("Failed to retrieve movies:")
        print(response.status_code, response.text)

def get_movie(movie_id):
    url = f"{BASE_URL}/movies/{movie_id}"
    response = requests.get(url)
    if response.status_code == 200:
        print("Movie retrieved successfully:")
        print(response.json())
    else:
        print("Failed to retrieve movie:")
        print(response.status_code, response.text)

# Example usage
if __name__ == "__main__":
    # Create a movie
    create_movie("Inception", 2010)

    # Retrieve all movies
    get_movies()

    # Retrieve a specific movie by ID (assuming ID 1 for this example)
    get_movie(1)

