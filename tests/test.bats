#!/usr/bin/env bats

set -eo pipefail
readonly build="$BATS_CWD/build/$(uname)/gh-release"

setup() {
  mkdir -p "$(dirname $BASH_SOURCE)/../release"
  touch "$(dirname $BASH_SOURCE)/../release/file.tgz"
}

teardown() {
  rm -rf "$(dirname $BASH_SOURCE)/../release"
}

@test "version" {
  cd "$(dirname $BASH_SOURCE)/.."

  run $build create progrium/gh-release test
  echo "status: $status"
  echo "output: $output"
  [[ "$status" -eq 0 ]]

  run $build destroy progrium/gh-release test
  echo "status: $status"
  echo "output: $output"
  [[ "$status" -eq 0 ]]

  cd - > /dev/null
}