all: lint build
lint:
	golangci-lint run
build:
	go build -o ./bin/golox
debug:
	go build -gcflags="all=-N -l" -o ./bin/golox-debug
	./bin/golox-debug
run: build
	./bin/golox