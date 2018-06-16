package chg

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSortChanges(t *testing.T) {
	v := &Version{
		Name: "1.0.0",
		Changes: []*ChangeList{
			{Type: Removed},
			{Type: Added},
			{Type: Fixed},
			{Type: Changed},
			{Type: Security},
			{Type: Deprecated},
		},
	}

	expected := []*ChangeList{
		{Type: Added},
		{Type: Changed},
		{Type: Deprecated},
		{Type: Fixed},
		{Type: Removed},
		{Type: Security},
	}

	v.SortChanges()

	assert.Equal(t, expected, v.Changes)
}

func TestChange(t *testing.T) {
	added := &ChangeList{Type: Added}
	removed := &ChangeList{Type: Removed}

	v := Version{
		Name: "1.0.0",
		Changes: []*ChangeList{
			added,
			removed,
		},
	}

	t.Run("change-exists", func(t *testing.T) {
		result := v.Change(Added)
		assert.Equal(t, added, result)
	})

	t.Run("change-does-not-exist", func(t *testing.T) {
		result := v.Change(Security)
		assert.Nil(t, result)
	})
}

func TestRenderTitle(t *testing.T) {
	t.Run("name-only", func(t *testing.T) {
		v := &Version{
			Name: "1.0.0",
		}
		expected := "## 1.0.0"
		var buf bytes.Buffer
		v.RenderTitle(&buf)
		result := string(buf.Bytes())
		assert.Equal(t, expected, result)
	})

	t.Run("date", func(t *testing.T) {
		v := &Version{
			Name: "1.0.0",
			Date: "2018-05-24",
		}
		expected := "## 1.0.0 - 2018-05-24"
		var buf bytes.Buffer
		v.RenderTitle(&buf)
		result := string(buf.Bytes())
		assert.Equal(t, expected, result)
	})

	t.Run("link", func(t *testing.T) {
		v := &Version{
			Name: "1.0.0",
			Link: "http://example.com/",
		}
		expected := "## [1.0.0]"
		var buf bytes.Buffer
		v.RenderTitle(&buf)
		result := string(buf.Bytes())
		assert.Equal(t, expected, result)
	})

	t.Run("yanked", func(t *testing.T) {
		v := &Version{
			Name:   "1.0.0",
			Yanked: true,
		}
		expected := "## 1.0.0 [YANKED]"
		var buf bytes.Buffer
		v.RenderTitle(&buf)
		result := string(buf.Bytes())
		assert.Equal(t, expected, result)
	})
}

func TestRenderChanges(t *testing.T) {
	changes := []*ChangeList{
		{
			Type: Added,
			Items: []*Item{
				{"Item 1"},
				{"Item 2"},
			},
		},
		{
			Type: Changed,
			Items: []*Item{
				{"Item A"},
				{"Item B"},
			},
		},
	}

	v := Version{Name: "1.0.0", Changes: changes}

	expected := `### Added
- Item 1
- Item 2

### Changed
- Item A
- Item B
`

	var buf bytes.Buffer
	v.RenderChanges(&buf)
	result := string(buf.Bytes())

	assert.Equal(t, expected, result)
}

func TestVersionRender(t *testing.T) {
	changes := []*ChangeList{
		{
			Type: Added,
			Items: []*Item{
				{"Item 1"},
				{"Item 2"},
			},
		},
		{
			Type: Changed,
			Items: []*Item{
				{"Item A"},
				{"Item B"},
			},
		},
	}

	v := Version{Name: "1.0.0", Changes: changes}

	expected := `## 1.0.0
### Added
- Item 1
- Item 2

### Changed
- Item A
- Item B
`

	var buf bytes.Buffer
	v.Render(&buf)
	result := string(buf.Bytes())

	assert.Equal(t, expected, result)
}
