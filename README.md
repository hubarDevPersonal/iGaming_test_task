# iGaming_test_task

## Description
This is a simple API that processes game data and returns the processed data. The API is built using go-chi and is dockerized.

## Setup Locally
1. Clone the repository
2. Run `docker compose up -d --build` in the root directory
3. Generate a sign header token using Secret key at the config
4. Use the generated token to make requests to the API "localhost:8080/api/v1/games-processor"