bulid:
	go build -o ./bin/golox
debug:
	go build -gcflags="all=-N -l" -o ./bin/golox-debug
	./bin/golox-debug