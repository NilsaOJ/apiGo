# Create another stage called "dev" that is based off of our "base" stage (so we have golang available to us)
FROM golang:1.17

# Create and change to the app directory.
WORKDIR /app/go-miamideas

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy local code to the container image.
COPY . .

# Build the binary.
RUN go build -v -o ./out/go-miamideas .

# Run the web service on container startup.
CMD ["./out/go-miamideas"]