# Start from base image
FROM golang:1.17-buster as builder

RUN mkdir /app

# Set the current working directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download && go mod verify

# Copy source from current directory to working directory
COPY . .

RUN go get github.com/githubnemo/CompileDaemon

EXPOSE 8000

ENTRYPOINT CompileDaemon --build="go build main.go" --command=./main