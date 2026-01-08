# To use make from windows user profile
# Run in the users folder with make -C <location> ...

build-http-image:
	docker build -t my-app:test-httpserver -f ./cmd/httpserver/Dockerfile .

build-grpc-image:
	docker build -t my-app:test-grpcserver -f ./cmd/grpcserver/Dockerfile .

build:
	golangci-lint run
	go test ./...

unit-tests:
	go test -short ./...