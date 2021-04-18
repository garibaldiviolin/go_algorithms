build:
	rm -Rf go_algorithms
	go build -o go_algorithms main.go

build-and-run: build
	./go_algorithms

format-style:
	go fmt ./...

run:
	go run main.go

test:
	go test ./...

test-coverage: test
	rm -Rf cover.out cover.html
	go test -v -coverprofile cover.out ./binary_search
	go tool cover -html=cover.out -o cover.html
