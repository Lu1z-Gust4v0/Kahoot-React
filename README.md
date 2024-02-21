# Kahoot React Clone

## Description 

This a kahoot clone made in NextJs with a Go backend. Realtime updates are handled using websockets

## Technologies 

- NextJs
- TailwindCSS
- Go 
- Fiber
- Postgresql
- Docker

## Running 

- Create a .env where .env.example exists

- Run Postgresql docker container
```
cd backend/
docker compose up db
```

- Run go server 
```
cd backend/
go run cmd/server.go
```

- Run frontend
```
npm run dev
```
