package chg

import (
	"bytes"
	"testing"
)

func TestItemRender(t *testing.T) {
	i := Item{"Item 1"}
	expected := "- Item 1\n"

	var buf bytes.Buffer
	i.Render(&buf)
	result := buf.String()

	if result != expected {
		t.Errorf("Item.Render failed, expected %s got %s", expected, result)
	}
}
