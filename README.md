# changelog

![Tests](https://github.com/cucumber/changelog/workflows/Tests/badge.svg)
[![Coverage Status](https://coveralls.io/repos/github/cucumber/changelog/badge.svg?branch=master)](https://coveralls.io/github/cucumber/changelog?branch=master)
[![Go Report Card](https://goreportcard.com/badge/github.com/cucumber/changelog)](https://goreportcard.com/report/github.com/cucumber/changelog)

(This is a fork of https://github.com/rcmachado/changelog)

`changelog` is a command-line application to read and manipulate
`CHANGELOG.md` files that follows the [keepachangelog.com][] spec.

It can normalize the file (`fmt`), create a release (`release`) and
show a specific version (`show`). See [Usage](#usage) for details.

## Table of Contents

- [Usage](#usage)
- [Installation](#installation)
  - [Linux and macOS](#linux-and-macos)
  - [Source](#source)
- [Commands](#commands)
  - [init](#init)
  - [fmt](#fmt)
  - [show](#show)
  - [release](#release)
- [Formatting](#formatting)
- [Contributing](#contributing)
- [License](#license)

## Usage

```bash
# Initialize a new CHANGELOG.md file
$ changelog init --compare-url=https://github.com/myorg/myrepo/compare/abcdef...1234
# Add a first message under "Added" section
$ changelog added "Initial commit"
# Create release
$ changelog release 0.1.0
```

## Installation

### Linux and macOS

The easiest way to install it is to download the [latest version][]
from GitHub releases.

There are precompiled binaries for macOS and Linux.

### Source

Clone the repository and build the executable:

```bash
make build
```

This will generate a `changelog` binary that can be copied to `/usr/local/bin`:

```bash
cp changelog /usr/local/bin
```

## Commands

```text
changelog manipulate and validate markdown changelog files following the keepachangelog.com specification.

Usage:
  changelog [command]

Available Commands:
  added       Add item under 'Added' section
  bundle      Bundles files containing unrelased changelog entries
  changed     Add item under 'Changed' section
  deprecated  Add item under 'Deprecated' section
  fixed       Add item under 'Fixed' section
  fmt         Reformat the change log file
  help        Help about any command
  init        Initializes a new changelog
  release     Change Unreleased to [version]
  removed     Add item under 'Removed' section
  security    Add item under 'Security' section
  show        Show changelog for [version]

Flags:
  -f, --filename string   Changelog file or '-' for stdin (default "CHANGELOG.md")
  -h, --help              help for changelog
  -o, --output string     Output file or '-' for stdout (default "-")

Use "changelog [command] --help" for more information about a command.
```

### init

Outputs a changelog with only preamble and Unreleased version to standard output. You can specify a filename using `--output/-o` flag:

```bash
$ touch CHANGELOG.md
$ changelog init -o CHANGELOG.md --compare-url https://github.com/cucumber/changelog/compare/abcdef...HEAD
Changelog file 'CHANGELOG.md' created.
```

### fmt

Normalize file format (see [Formatting](#formatting) for the specific
transformation applied):

```bash
changelog fmt
```

### show

Show what will be in the next release:

```bash
changelog show Unreleased
```

Show the change log for a specific version:

```bash
changelog show 1.2.3
```

### release

Create a new release:

```bash
changelog release 1.2.4
```

### Formatting

`fmt` command normalizes the changelog file. The idea is to always have
the same output, no matter how messy the file is. Right now it doesn't
do much, but the plan is to evolve it as a kind of `go fmt` for
changelogs.

Currently, the following transformations are applied:

- Sections are sorted (eg. Added, Changed, etc)
- Version links are put at the bottom of the file
- List bullet is always `-`

## Contributing

Feel free to fork and submit a PR. You can also take a look, at the [Issues][] tab to see some ideas.

## License

Licensed under MIT. See [LICENSE][] file for details.

[keepachangelog.com]: https://keepachangelog.com/
[LICENSE]: ./LICENSE
[Issues]: https://github.com/cucumber/changelog
[latest version]: https://github.com/cucumber/changelog/releases/latest
