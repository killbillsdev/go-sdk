.PHONY: test

test:
	cd go_sdk && go test
install:
	go get -v ./...