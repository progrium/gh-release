
install:
	install gh-release /usr/local/bin

release:
	rm -rf release
	mkdir release
	echo "$(VERSION)" > release/version
	cp ./gh-release release/gh-release
	
	