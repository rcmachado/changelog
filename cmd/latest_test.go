package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLatestCmdShowsLatestReleasedVersion(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/show-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	expected := `1.0.0
`

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: out,
	}

	cmd := newLatestCmd(iostreams)
	_, err = cmd.ExecuteC()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(out.Bytes()))
}

func TestLatestCmdError(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/empty-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: out,
	}

	cmd := newLatestCmd(iostreams)
	_, err = cmd.ExecuteC()

	assert.Error(t, err)
	expected := "There are no versions in the changelog yet"
	assert.EqualError(t, err, expected)
}

func TestLatestCmdErrorWhenNoReleasedVersions(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/minimal-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: out,
	}

	cmd := newLatestCmd(iostreams)
	_, err = cmd.ExecuteC()

	assert.Error(t, err)
	expected := "There are no released versions in the changelog yet"
	assert.EqualError(t, err, expected)
}
