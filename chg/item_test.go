package chg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestItemRender(t *testing.T) {
	i := Item{"Item 1"}
	expected := "- Item 1\n"

	var buf bytes.Buffer
	i.Render(&buf)
	result := buf.String()

	assert.Equal(t, expected, result)
}
