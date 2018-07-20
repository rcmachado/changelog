# changelog

[![CircleCI](https://circleci.com/gh/rcmachado/changelog.svg?style=svg)](https://circleci.com/gh/rcmachado/changelog)
[![Go Report Card](https://goreportcard.com/badge/github.com/rcmachado/changelog)](https://goreportcard.com/report/github.com/rcmachado/changelog)

`changelog` is a command-line application to read and manipulate
`CHANGELOG.md` files that follows the [keepachangelog.com][] spec.

## Installation

Download and build the executable:

```bash
$ go get -u github.com/rcmachado/changelog
$ cd $GOPATH/github.com/rcmachado/changelog
$ make build
```

## Usage

Normalize file format (see [Formatting](#formatting) for the specific
transformation applied):

```bash
$ changelog fmt
```

Show what will be in the next release:

```bash
$ changelog show Unreleased
```

Show the change log for a specific version:

```bash
$ changelog show 1.2.3
```

### Formatting

`fmt` command normalizes the changelog file. The idea is to always have
the same output, no matter how messy the file is. Right now it doesn't
do much, but the plan is to evolve it as a kind of `go fmt` for
changelogs.

The transformations applied are:

- Sections are sorted (eg. Added, Changed, etc)
- Version links are put at the bottom of the file
- List bullet is always `-`

## Contributing

To contribute, you can fork the repository and submit a Pull Request.

### What needs to be done?

A lot of things! Take a look at the issues to find something.

### I found a bug

Could you please [fill an issue][] explaining what caused the issue and
what did you expect to happen instead?

## License

It's released under MIT license. See [LICENSE][] file for details.

[keepachangelog.com]: https://keepachangelog.com/
[LICENSE]: ./LICENSE
[fill an issue]: https://github.com/rcmachado/changelog/issues
