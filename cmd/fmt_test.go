package cmd

import (
	"bytes"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFmtCmd(t *testing.T) {
	changelog := `# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).
	
## [Unreleased]
### Changed
- Out of order entries
- Another item here

### Added
- Something else

## [0.1.0] - 2018-06-17

### Added

- Command A
- Command B
- Command C

[Unreleased]: https://github.com/rcmachado/changelog/compare/0.2.0...HEAD
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0`

	expected := `# Changelog

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

[Unreleased]: https://github.com/rcmachado/changelog/compare/0.2.0...HEAD
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0
`

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  strings.NewReader(changelog),
		Out: out,
	}

	fmt := newFmtCmd(iostreams)
	_, err := fmt.ExecuteC()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(out.Bytes()))
}

func TestFmtCmdJson(t *testing.T) {
	changelog := `# Changelog
Hello
## [Unreleased]
### Changed
- Out of order entries

### Added
- Something else

## [0.1.0] - 2018-06-17

### Added

- Command A

[Unreleased]: https://github.com/rcmachado/changelog/compare/0.2.0...HEAD
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0`

	expected := `{
  "preamble": "Hello",
  "versions": [
    {
      "name": "Unreleased",
      "date": "",
      "link": "https://github.com/rcmachado/changelog/compare/0.2.0...HEAD",
      "yanked": false,
      "changes": [
        {
          "type": 2,
          "items": [
            {
              "description": "Out of order entries"
            }
          ]
        },
        {
          "type": 1,
          "items": [
            {
              "description": "Something else"
            }
          ]
        }
      ]
    },
    {
      "name": "0.1.0",
      "date": "2018-06-17",
      "link": "https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0",
      "yanked": false,
      "changes": [
        {
          "type": 1,
          "items": [
            {
              "description": "Command A"
            }
          ]
        }
      ]
    }
  ]
}
`

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  strings.NewReader(changelog),
		Out: out,
	}

	fmt := newFmtCmd(iostreams)
	fmt.SetArgs([]string{"--json"})
	_, err := fmt.ExecuteC()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(out.Bytes()))
}
