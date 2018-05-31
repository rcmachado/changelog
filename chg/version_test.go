package chg

import (
	"bytes"
	"reflect"
	"testing"
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

	if !reflect.DeepEqual(v.Changes, expected) {
		t.Error("SortChanges should sort Changes properly")
	}
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
		if result != added {
			t.Errorf("Search for change, expected %s got %s", Added, result)
		}
	})

	t.Run("change-does-not-exist", func(t *testing.T) {
		result := v.Change(Security)
		if result != nil {
			t.Errorf("Search for unknown change, expected nil got %s", result)
		}
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
		if expected != result {
			t.Errorf("RenderTitle should render version only, expected %s, got %s", expected, result)
		}
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
		if expected != result {
			t.Errorf("RenderTitle should render the date, expected %s, got %s", expected, result)
		}
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
		if expected != result {
			t.Errorf("RenderTitle should render link, expected %s, got %s", expected, result)
		}
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
		if expected != result {
			t.Errorf("RenderTitle should render yanked versions, expected %s, got %s", expected, result)
		}
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

	if result != expected {
		t.Errorf("RenderChanges fail, expected %s got %s", expected, result)
	}
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

	if result != expected {
		t.Errorf("RenderChanges fail, expected %s got %s", expected, result)
	}
}
