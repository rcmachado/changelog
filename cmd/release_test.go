package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseCmd(t *testing.T) {
	changelog := `# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Something else

### Changed
- Out of order entries
- Another item here

## [0.1.0] - 2018-06-17
### Added
- Command A
- Command B
- Command C

[Unreleased]: https://github.com/rcmachado/changelog/compare/0.1.0...HEAD
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0
`

	expected := `# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.2.0] - 2018-06-18
### Added
- Something else

### Changed
- Out of order entries
- Another item here

## [0.1.0] - 2018-06-17
### Added
- Command A
- Command B
- Command C

[Unreleased]: https://github.com/rcmachado/changelog/compare/0.2.0...HEAD
[0.2.0]: https://github.com/rcmachado/changelog/compare/0.1.0...0.2.0
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0
`

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  strings.NewReader(changelog),
		Out: out,
	}

	release := newReleaseCmd(iostreams)
	release.SetArgs([]string{"0.2.0", "--release-date", "2018-06-18"})
	_, err := release.ExecuteC()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(out.Bytes()))
}
