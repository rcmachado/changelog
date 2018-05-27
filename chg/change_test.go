package chg

import (
	"bytes"
	"testing"
)

func TestNewChange(t *testing.T) {
	t.Run("type=added", func(t *testing.T) {
		result := NewChange("Added")
		if result == nil || result.Type != Added {
			t.Errorf("NewChange failed expected Added got %s", result)
		}
	})

	t.Run("type=changed", func(t *testing.T) {
		result := NewChange("Changed")
		if result == nil || result.Type != Changed {
			t.Errorf("NewChange failed expected Changed got %s", result)
		}
	})

	t.Run("type=deprecated", func(t *testing.T) {
		result := NewChange("Deprecated")
		if result == nil || result.Type != Deprecated {
			t.Errorf("NewChange failed expected Deprecated got %s", result)
		}
	})

	t.Run("type=fixed", func(t *testing.T) {
		result := NewChange("Fixed")
		if result == nil || result.Type != Fixed {
			t.Errorf("NewChange failed expected Fixed got %s", result)
		}
	})

	t.Run("type=removed", func(t *testing.T) {
		result := NewChange("Removed")
		if result == nil || result.Type != Removed {
			t.Errorf("NewChange failed expected Removed got %s", result)
		}
	})

	t.Run("type=security", func(t *testing.T) {
		result := NewChange("Security")
		if result == nil || result.Type != Security {
			t.Errorf("NewChange failed expected Security got %s", result)
		}
	})

	t.Run("type=unknown", func(t *testing.T) {
		result := NewChange("unknown")
		if result != nil {
			t.Errorf("NewChange failed expected nil got %s", result)
		}
	})
}

func TestChangeRender(t *testing.T) {
	c := Change{Type: Added, Content: "something"}

	expected := "### Added\nsomething"

	var buf bytes.Buffer
	c.Render(&buf)
	result := string(buf.Bytes())
	if result != expected {
		t.Errorf("Render failed, expected %s got %s", expected, result)
	}
}
