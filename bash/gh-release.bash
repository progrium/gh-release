
readonly ref_endpoint="${GITHUB_API_URL:-https://api.github.com}/repos/%s/git/refs/tags/%s"
readonly release_endpoint="${GITHUB_API_URL:-https://api.github.com}/repos/%s/releases"
readonly release_json='{"tag_name": "v%s", "name": "%s", "target_commitish": "%s"}'

release-create() {
	declare reponame="$1" version="${2#v}" branch="${3:-master}" name="$4"
	local release="$(printf "$release_json" "$version" "$name" "$branch")"
	local release_url="$(printf "$release_endpoint" "$reponame")"
	echo "Creating release v$version from branch $branch ..."
	upload_url="$(curl -s -H "Authorization: token $GITHUB_ACCESS_TOKEN" -d "$release" "$release_url" | upload-url)"
	for asset in $(ls -A release); do
		local name="$(basename $asset)"
		echo "Uploading $name ..."
		curl -X POST -H "Content-Type: $(mimetype $name)" -H "Authorization: token $GITHUB_ACCESS_TOKEN" --data-binary "@release/$asset" \
			"$upload_url=$name" > /dev/null
	done
}

release-destroy() {
	declare reponame="$1" version="$2"
	local release_url="$(printf "$release_endpoint" "$reponame")"

    [[ "$version" == [0-9]* ]] && version="v$version"

	release_id="$(curl -s -H "Authorization: token $GITHUB_ACCESS_TOKEN" "$release_url" | release-id-from-tagname "$version")"
	echo "Deleting release..."
	curl -s -H "Authorization: token $GITHUB_ACCESS_TOKEN" -X DELETE "$release_url/$release_id"
	echo "Deleting tag..."
	tag_url="$(printf "$ref_endpoint" "$reponame" "$version")"
	curl -s -H "Authorization: token $GITHUB_ACCESS_TOKEN" -X DELETE "$tag_url"
}

usage() {
	echo "Usage: gh-release [-v] subcommand"
	echo
	echo "Subcommands:"
	echo "  create <reponame> <version> [branch] [name]"
	echo "  destroy <reponame> <version>"
	echo "  checksums <algorithm>"
	echo
}

release-checksums() {
	declare alg="$1"
	echo "Writing $alg checksum files..."
	for asset in $(ls -A release); do
		cat "release/$asset" | checksum "$alg" > "release/${asset}.$alg"
	done
}

main() {
	set -eo pipefail; [[ "$TRACE" ]] && set -x
	case "$1" in
		create)		shift; release-create "$@";;
		destroy)	shift; release-destroy "$@";;
		checksums)	shift; release-checksums "$@";;
		-v)			echo "$VERSION";;
		*)			usage;;
	esac
}
