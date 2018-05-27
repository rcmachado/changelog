package chg

import (
	"bytes"
	"testing"
)

func TestNewChangeList(t *testing.T) {
	t.Run("type=added", func(t *testing.T) {
		result := NewChangeList("Added")
		if result == nil || result.Type != Added {
			t.Errorf("NewChangeList failed expected Added got %s", result)
		}
	})

	t.Run("type=changed", func(t *testing.T) {
		result := NewChangeList("Changed")
		if result == nil || result.Type != Changed {
			t.Errorf("NewChangeList failed expected Changed got %s", result)
		}
	})

	t.Run("type=deprecated", func(t *testing.T) {
		result := NewChangeList("Deprecated")
		if result == nil || result.Type != Deprecated {
			t.Errorf("NewChangeList failed expected Deprecated got %s", result)
		}
	})

	t.Run("type=fixed", func(t *testing.T) {
		result := NewChangeList("Fixed")
		if result == nil || result.Type != Fixed {
			t.Errorf("NewChangeList failed expected Fixed got %s", result)
		}
	})

	t.Run("type=removed", func(t *testing.T) {
		result := NewChangeList("Removed")
		if result == nil || result.Type != Removed {
			t.Errorf("NewChangeList failed expected Removed got %s", result)
		}
	})

	t.Run("type=security", func(t *testing.T) {
		result := NewChangeList("Security")
		if result == nil || result.Type != Security {
			t.Errorf("NewChangeList failed expected Security got %s", result)
		}
	})

	t.Run("type=unknown", func(t *testing.T) {
		result := NewChangeList("unknown")
		if result != nil {
			t.Errorf("NewChangeList failed expected nil got %s", result)
		}
	})
}

func TestChangeRender(t *testing.T) {
	c := ChangeList{Type: Added, Content: "something"}

	expected := "### Added\nsomething"

	var buf bytes.Buffer
	c.Render(&buf)
	result := string(buf.Bytes())
	if result != expected {
		t.Errorf("Render failed, expected %s got %s", expected, result)
	}
}
