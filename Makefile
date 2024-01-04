version ?= 0.0.3

publish: install build test reindex

reindex:
	@echo "Publishing new version ! ..."
	git tag v$(version) && git push origin v$(version) && GOPROXY=proxy.golang.org go list -m github.com/killbillsdev/go_sdk@v$(version)
	@echo "Done"

install:
	@echo "Installing..."
	cd go_sdk && go get .
	@echo "Done"

build:
	@echo "Building..."
	go build -v ./...
	@echo "Done"

test:
	@echo "Testing..."
	cd go_sdk && go test
	@echo "Done"

.DEFAULT_GOAL := install
.PHONY: build install publish
