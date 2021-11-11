# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Fixed
- Show command was broken ([#4](https://github.com/cucumber/changelog/issues/4))

## [0.8.0] - 2021-11-11
### Added
- Add `--json` flag to `fmt` and `show` commands. ([#1](https://github.com/cucumber/changelog/pull/1))
- Add `--tag-format` to `release` command. ([#2](https://github.com/cucumber/changelog/pull/2))

## [0.7.0] - 2020-07-03
### Changed
- Install git and openssh on docker image

### Fixed
- Generate compressed archives for each release

## [0.6.0] - 2020-05-28
### Changed
- Change docker base image from `scratch` to `debian:bullseye`

## [0.5.0] - 2020-05-26
### Added
- Publish docker image to [rcmachado/changelog](https://hub.docker.com/r/rcmachado/changelog) when releasing

## [0.4.2] - 2020-03-31
### Fixed
- Compare URL now parses versions with dots correctly

## [0.4.1] - 2020-02-27
### Fixed
- Wrong version for blackfriday dependency

## [0.4.0] - 2020-02-27
### Fixed
- `release` command when there is only `Unreleased` and no compare link was provided

### Removed
- `bundle` command

## [0.3.0] - 2020-01-08
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

[Unreleased]: https://github.com/rcmachado/changelog/compare/0.8.0...HEAD
[0.8.0]: https://github.com/rcmachado/changelog/compare/0.7.0...0.8.0
[0.7.0]: https://github.com/rcmachado/changelog/compare/0.6.0...0.7.0
[0.6.0]: https://github.com/rcmachado/changelog/compare/0.5.0...0.6.0
[0.5.0]: https://github.com/rcmachado/changelog/compare/0.4.2...0.5.0
[0.4.2]: https://github.com/rcmachado/changelog/compare/0.4.1...0.4.2
[0.4.1]: https://github.com/rcmachado/changelog/compare/0.4.0...0.4.1
[0.4.0]: https://github.com/rcmachado/changelog/compare/0.3.0...0.4.0
[0.3.0]: https://github.com/rcmachado/changelog/compare/0.2.0...0.3.0
[0.2.0]: https://github.com/rcmachado/changelog/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0
