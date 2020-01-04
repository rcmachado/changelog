# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Changed
- Replaced [dep](https://github.com/golang/dep) with [Go Modules](https://blog.golang.org/using-go-modules)
- Replaced [gometalinter.v2](https://github.com/alecthomas/gometalinter) with [golang-ci](https://github.com/golangci/golangci-lint)
- Updated [stretchr/testify](github.com/stretchr/testify) to v1.4.0
- Updated [spf13/cobra](github.com/spf13/cobra) to v0.0.5

## [0.2.0] - 2018-07-20
### Added
- `-o/--output` option

### Changed
- `bump` command renamed to `release`
- Use `dep` to handle dependencies

### Fixed
- Merge duplicated sections when parsing changelog file

## [0.1.0] - 2018-06-17
### Added
- `bump` command to increment a release
- `fmt` command to reformat changelog following the spec
- `show` command to show a specific version

[Unreleased]: https://github.com/rcmachado/changelog/compare/0.2.0...HEAD
[0.2.0]: https://github.com/rcmachado/changelog/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0
