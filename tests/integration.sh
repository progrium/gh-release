
readonly build="$(dirname $BASH_SOURCE)/../build/$(uname)/gh-release"

testWorks() {
	cd "$(dirname $BASH_SOURCE)/.."
	rm -rf release && mkdir release
	touch release/file.tgz
	$build create progrium/gh-release test
	$build destroy progrium/gh-release test
	cd - > /dev/null
}