package cmd

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInitCmd(t *testing.T) {
	expected := `# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- First commit

[Unreleased]: https://github.com/rcmachado/changelog/compare/abcdef...HEAD
`

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  nil,
		Out: out,
	}

	cmd := newInitCmd(iostreams)
	cmd.SetArgs([]string{"--compare-url", "https://github.com/rcmachado/changelog/compare/abcdef...HEAD"})
	_, err := cmd.ExecuteC()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(out.Bytes()))
}

func TestInitCmdRequiresCompareURL(t *testing.T) {
	iostreams := &IOStreams{
		In:  nil,
		Out: new(bytes.Buffer),
	}

	cmd := newInitCmd(iostreams)
	_, err := cmd.ExecuteC()

	assert.Error(t, err)
}
