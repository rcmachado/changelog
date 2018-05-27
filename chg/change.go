package chg

//go:generate stringer -type=ChangeType

import (
	"fmt"
	"io"
	"strings"
)

// Change groups the changes by type
// Valid change types are "Added", "Changed", "Deprecated", "Fixed",
// "Removed" and "Security"
type Change struct {
	Type    ChangeType
	Content string
}

// ChangeType is the type of the changes
type ChangeType int

// Change types
const (
	Added ChangeType = iota
	Changed
	Deprecated
	Fixed
	Removed
	Security
)

// NewChange creates a Change struct based on the informed type
func NewChange(ct string) *Change {
	change := &Change{}
	switch strings.ToLower(ct) {
	case "added":
		change.Type = Added
	case "changed":
		change.Type = Changed
	case "deprecated":
		change.Type = Deprecated
	case "fixed":
		change.Type = Fixed
	case "removed":
		change.Type = Removed
	case "security":
		change.Type = Security
	default:
		return nil
	}
	return change
}

// Render builds the representation of Change
func (c *Change) Render(w io.Writer) {
	io.WriteString(w, fmt.Sprintf("### %s\n", c.Type))
	io.WriteString(w, c.Content)
}
