NAME=gh-release
HARDWARE=$(shell uname -m)
BRANCH=$(shell git rev-parse --abbrev-ref HEAD)
VERSION=2.0.1

build:
	go-bindata bash
	mkdir -p build/Linux && GOOS=linux go build -o build/Linux/$(NAME)
	mkdir -p build/Darwin && GOOS=darwin go build -o build/Darwin/$(NAME)

deps:
	go get -u github.com/jteeuwen/go-bindata/...
	go get -u github.com/progrium/gh-release/...
	go get || true

test: build
	tests/shunit2 tests/*.sh

release: build
	rm -rf release && mkdir release
	tar -zcf release/$(NAME)_$(VERSION)_linux_$(HARDWARE).tgz -C build/Linux $(NAME)
	tar -zcf release/$(NAME)_$(VERSION)_darwin_$(HARDWARE).tgz -C build/Darwin $(NAME)
	build/$(shell uname)/gh-release create progrium/$(NAME) $(VERSION) $(BRANCH)

.PHONY: release build deps test