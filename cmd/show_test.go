package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestShowCmd(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/show-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	expected := `### Added
- Item 1
- Item 2

### Changed
- Item 3
`

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: out,
	}

	cmd := newShowCmd(iostreams)
	cmd.SetArgs([]string{"1.0.0"})
	_, err = cmd.ExecuteC()

	assert.Nil(t, err)
	assert.Equal(t, expected, string(out.Bytes()))
}

func TestShowCmdUnknownVersion(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/show-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: out,
	}

	cmd := newShowCmd(iostreams)
	cmd.SetArgs([]string{"9.9.9"})
	_, err = cmd.ExecuteC()

	assert.NotNil(t, err)
}
