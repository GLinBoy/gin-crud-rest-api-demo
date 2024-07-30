# Gin CRUD REST API demo

## Setup database

There is a .env file, contains database data and cerdentials. don't forget to check it and update it if you need it.

## cURL requests

- Create Movie

  ```bash
    curl --location 'localhost:8080/movies' \
    --header 'Content-Type: application/json' \
    --data '{
    "name": "Top Gun",
    "description": "Top Gun: Maverick is a 2022 American action drama film directed by Joseph Kosinski and written by Ehren Kruger, Eric Warren Singer, and Christopher McQuarrie from a story by Peter Craig and Justin Marks."
    }'
  ```

- Get Movies

  ```bash
  curl --location 'localhost:8080/movies'
  ```

- Get a Movie

  ```bash
  curl --location 'localhost:8080/movies/696b1743-0efa-4b5c-a254-d71c2d3eede0'
  ```

- Update Movie

  ```bash
  curl --location --request PUT 'localhost:8080/movies/696b1743-0efa-4b5c-a254-d71c2d3eede0' \
  --header 'Content-Type: application/json' \
  --data '{
      "name": "Top Gun",
      "description": "Top Gun: Maverick is a 2022 American action drama film directed by Joseph Kosinski and written by Ehren Kruger, Eric Warren Singer, and Christopher McQuarrie from a story by Peter Craig and Justin Marks."
  }'
  ```

- Delete Movie

  ```bash
  curl --location --request DELETE 'localhost:8080/movies/696b1743-0efa-4b5c-a254-d71c2d3eede0'
  ```
