REPO=progrium/gh-release
VERSION=1.0.0

install:
	install gh-release /usr/local/bin

release:
	rm -rf release
	mkdir release
	echo "$(VERSION)" > release/version
	echo "$(VERSION)" > release/name
	echo "$(REPO)" > release/repo
	cp ./gh-release release/gh-release
	./gh-release gh-release

.PHONY: release