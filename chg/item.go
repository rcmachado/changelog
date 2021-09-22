package chg

import (
	"fmt"
	"io"
)

// Item holds the change itself
type Item struct {
	Description string `json:"description"`
}

// Render renders the change as a list item
func (i *Item) Render(w io.Writer) {
	io.WriteString(w, fmt.Sprintf("- %s\n", i.Description))
}
