FROM golang:latest

WORKDIR /app
RUN go get -u github.com/go-redis/redis
COPY . /app
EXPOSE 5000
CMD [ "go", "run", "shortener.go", "server.go" ]