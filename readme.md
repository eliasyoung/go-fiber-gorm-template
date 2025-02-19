# Golang Project Template
- go-fiber
- gorm

## Initialize DB
Check and edit docker compose file in ./docker/

Run following command to create a db container:
```bash
docker compose -f db-docker-compose.yaml up
```

There's also an example file if you need initial db container with another port instead of 5432

## Deployment in Docker
### 1. Build the api server with Dockerfile
```bash
docker build -t my-go-server-image .
```
### 2. Push Image to your host or dockerhub
### 3. Config
Edit docker-compose.yaml and .env(copy from .env.example if not exist) under ./docker
### 4. Start container with docker compose
```bash
docker compose up
```

## Auto rebuild in development with air
Make sure you have air installed fisrt.
Copy and edit content of .air.toml.example to .air.toml
```bash
go install github.com/cosmtrek/air@latest
```
Then run following command to start development server:
```bash
air
```
