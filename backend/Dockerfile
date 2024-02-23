FROM golang:1.21.7-alpine3.19 AS base

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . . 

RUN CGO_ENABLE=0 GOOS=linux go build -o server cmd/server.go

EXPOSE 8000 

CMD ["./server"]
