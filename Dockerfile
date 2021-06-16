FROM golang:alpine
COPY main.go go.mod go.sum /app/
WORKDIR /app
CMD go run main.go