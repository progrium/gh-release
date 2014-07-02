# gh-release

Bash script for creating and uploading files to Github Releases.

## Getting gh-release

You can download into a PATH directory from a stable release and chmod:

	$ wget -O /usr/local/bin/gh-release https://github.com/progrium/gh-release/releases/download/v1.0.0/gh-release
	$ chmod +x /usr/local/bin/gh-release

Or you can clone the repo and have it install to `/usr/local/bin/gh-release`:

	$ make install

## Dependencies

It currently depends on `jq`, so just `apt-get` or `brew` install `jq` if you don't have it.

Also, [get a GitHub personal access token](https://help.github.com/articles/creating-an-access-token-for-command-line-use) and put it in the environment variable `GITHUB_ACCESS_TOKEN`. Ideally you should export this in your `.profile` file. 

## Using gh-release

First, when you run gh-release, it's going to create a tag and release from your repo's master branch. So be sure you've pushed to master before running gh-release.

Next, gh-release assumes you'll be calling from a Makefile that also created a `release` directory. In this directory you put your files to upload. You also put several files to help gh-release:

 * release/repo: The user and repo name to use for the release. Example: "progrium/gh-release"
 * release/version: The version to use for release. "1.0.1" will result in a release/tag called "v1.0.1"
 * release/name: Optional name to use of the release. Without this, Github uses the version as the name

Once you have a `release` directory with at least `repo` and `version` files and the files to upload, you can run `gh-release` from the parent directory of `release` (ie, where Makefile runs). 

By default, it will look for `*.tgz` files to upload. But you can change this by passing an argument like `gh-release *.zip` or `gh-release myfile`, where `myfile` is a single file in `release` you want to upload. 

See this project's Makefile for an example.

## Sponsor

This project was made possible thanks to [DigitalOcean](http://digitalocean.com).

## License

BSD