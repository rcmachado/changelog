package chg

import (
	"fmt"
	"io"
)

// Item holds the change itself
type Item struct {
	Description string
}

// Render rendes the change as a list item
func (i *Item) Render(w io.Writer) {
	io.WriteString(w, fmt.Sprintf("- %s\n", i.Description))
}
