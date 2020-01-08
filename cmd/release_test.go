package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReleaseCmd(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/minimal-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	expected := `# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](http://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](http://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2018-06-18
### Added
- Item 1

[Unreleased]: https://example.com/0.1.0/HEAD
[0.1.0]: https://github.com/rcmachado/changelog/compare/ae761ff...0.1.0
`

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: out,
	}

	release := newReleaseCmd(iostreams)
	release.SetArgs([]string{"0.1.0", "--release-date", "2018-06-18", "--compare-url", "https://example.com/<prev>/<next>"})
	_, err = release.ExecuteC()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(out.Bytes()))
}

func TestReleaseCmdError(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/minimal-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: new(bytes.Buffer),
	}

	release := newReleaseCmd(iostreams)
	// Missing --compare-url, as the autodetect won't work for the minimal changelog
	release.SetArgs([]string{"0.1.0", "--release-date", "2018-06-18"})
	_, err = release.ExecuteC()

	assert.Error(t, err)
}
