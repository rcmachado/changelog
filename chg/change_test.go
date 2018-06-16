package chg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewChangeList(t *testing.T) {
	t.Run("type=added", func(t *testing.T) {
		result := NewChangeList("Added")

		assert.NotNil(t, result)
		assert.Equal(t, Added, result.Type)
	})

	t.Run("type=changed", func(t *testing.T) {
		result := NewChangeList("Changed")

		assert.NotNil(t, result)
		assert.Equal(t, Changed, result.Type)
	})

	t.Run("type=deprecated", func(t *testing.T) {
		result := NewChangeList("Deprecated")

		assert.NotNil(t, result)
		assert.Equal(t, Deprecated, result.Type)
	})

	t.Run("type=fixed", func(t *testing.T) {
		result := NewChangeList("Fixed")

		assert.NotNil(t, result)
		assert.Equal(t, Fixed, result.Type)
	})

	t.Run("type=removed", func(t *testing.T) {
		result := NewChangeList("Removed")

		assert.NotNil(t, result)
		assert.Equal(t, Removed, result.Type)
	})

	t.Run("type=security", func(t *testing.T) {
		result := NewChangeList("Security")

		assert.NotNil(t, result)
		assert.Equal(t, Security, result.Type)
	})

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
