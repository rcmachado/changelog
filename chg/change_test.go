package chg

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChangeList(t *testing.T) {
	var testData = []struct {
		input    string
		expected ChangeType
	}{
		{input: "Added", expected: Added},
		{input: "Changed", expected: Changed},
		{input: "Deprecated", expected: Deprecated},
		{input: "Fixed", expected: Fixed},
		{input: "Removed", expected: Removed},
		{input: "Security", expected: Security},
	}

	for _, tt := range testData {
		t.Run(fmt.Sprintf("type=%s", tt.input), func(t *testing.T) {
			result := NewChangeList(tt.input)

			assert.NotNil(t, result)
			assert.Equal(t, tt.expected, result.Type)
		})
	}

	t.Run("type=unknown", func(t *testing.T) {
		result := NewChangeList("unknown")

		assert.Nil(t, result)
	})
}

func TestChangeListRenderItems(t *testing.T) {
	c := ChangeList{
		Items: []*Item{
			{"Item 1"},
			{"Item 2"},
			{"Item 3"},
		},
	}
	expected := `- Item 1
- Item 2
- Item 3
`

	var buf bytes.Buffer
	c.RenderItems(&buf)
	result := buf.String()

	assert.Equal(t, expected, result)
}

func TestChangeRender(t *testing.T) {
	c := ChangeList{
		Type: Added,
		Items: []*Item{
			{"something"},
		},
	}

	expected := "### Added\n- something\n"

	var buf bytes.Buffer
	c.Render(&buf)
	result := string(buf.Bytes())

	assert.Equal(t, expected, result)
}
