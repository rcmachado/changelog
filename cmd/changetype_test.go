package cmd

import (
	"bytes"
	"io/ioutil"
	"testing"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/stretchr/testify/assert"
)

func TestNewChangeTypeCmd(t *testing.T) {
	changelog, err := ioutil.ReadFile("testdata/minimal-changelog.md")
	if err != nil {
		t.Fatal(err)
	}

	out := new(bytes.Buffer)
	iostreams := &IOStreams{
		In:  bytes.NewBuffer(changelog),
		Out: out,
	}

	cmd := newChangeTypeCmd(iostreams, chg.Security)
	cmd.SetArgs([]string{"some", "release", "item"})
	_, err = cmd.ExecuteC()

	assert.Nil(t, err)

	c := parser.Parse(out)
	v := c.Version("Unreleased")

	secChange := v.Change(chg.Security)
	assert.NotNil(t, secChange)
	assert.NotEmpty(t, secChange.Items)
	assert.Equal(t, "some release item", secChange.Items[0].Description)
}

func TestNewChangeTypeCmds(t *testing.T) {
	iostreams := &IOStreams{
		In:  new(bytes.Buffer),
		Out: new(bytes.Buffer),
	}

	allCmds := newChangeTypeCmds(iostreams)
	assert.Len(t, allCmds, 6)
}
