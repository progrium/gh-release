# gh-release

Utility for creating, deleting, and uploading files to Github Releases.

[![Circle CI](https://circleci.com/gh/progrium/gh-release.png?style=shield)](https://circleci.com/gh/progrium/gh-release)

## Getting gh-release

Download and uncompress the appropriate binary from [releases](https://github.com/progrium/gh-release/releases).

## Using gh-release

You need to have a [Github personal access token](https://help.github.com/articles/creating-an-access-token-for-command-line-use) defined in your environment as `GITHUB_ACCESS_TOKEN`.

	$ gh-release 
	Usage: gh-release create|destroy <reponame> <version> [branch] [name]

#### Creating a release with assets

Put any assets you want to upload with your release into a `release` directory. Then call `gh-release`. Here is an example:

	$ gh-release create progrium/gh-release 1.0.0

This will create the release then upload any files in the `release` directory. Optionally, you can pass the branch to tag the release from, as well as a name for the release. 

See this project's Makefile for a full example of using it in a Makefile.

#### Destroying a release

You can destroy a release by the version number you used to create the release:

	$ gh-release destroy progrium/gh-release 1.0.0

This destroys the release and its assets, as well as the tag created for the release.

## Sponsor

This project was made possible thanks to [DigitalOcean](http://digitalocean.com).

## License

BSD