BINARY_NAME=hwapi

build:
	@echo "Building..."
	env CGO_ENABLED=0 go build -ldflags="-s -w" -o $(BINARY_NAME) -v
	@echo "Done."

run: build
	@echo "Starting..."
	./${BINARY_NAME} &
	@echo "Started!"

clean:
	@echo "Cleaning..."
	@go clean
	@rm ${BINARY_NAME}
	@echo "Cleaned!"

stop:
	@echo "Stopping..."
	@-pkill -SIGTERM -f "./${BINARY_NAME}"
	@echo "Stopped!"

restart: stop start

test:
	go test -v ./...
