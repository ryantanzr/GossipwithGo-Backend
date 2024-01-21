.PHONY: build run test clean

BINARY_NAME=main
CMD_PATH=./cmd/server

build:
    @echo "Building..."
    @go build -o $(BINARY_NAME) $(CMD_PATH)/main.go

run: build
    @echo "Running..."
    @./$(BINARY_NAME)

test: 
    @echo "Testing..."
    @go test -v ./...

clean:
    @echo "Cleaning..."
    @rm -f $(BINARY_NAME)