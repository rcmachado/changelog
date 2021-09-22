package chg

//go:generate stringer -type=ChangeType

import (
	"fmt"
	"io"
	"strings"
)

// ChangeList groups the changes by type
// Valid change types are "Added", "Changed", "Deprecated", "Fixed",
// "Removed" and "Security"
type ChangeList struct {
	Type  ChangeType `json:"type"`
	Items []*Item    `json:"items"`
}

// ChangeType is the type of the changes
type ChangeType int

// Change types
const (
	Unknown ChangeType = iota
	Added
	Changed
	Deprecated
	Fixed
	Removed
	Security
)

// ChangeTypeFromString creates a type based on its string name
func ChangeTypeFromString(ct string) ChangeType {
	switch strings.ToLower(ct) {
	case "added":
		return Added
	case "changed":
		return Changed
	case "deprecated":
		return Deprecated
	case "fixed":
		return Fixed
	case "removed":
		return Removed
	case "security":
		return Security
	default:
		return Unknown
	}
}

// NewChangeList creates a ChangeList struct based on the informed type
func NewChangeList(ct string) *ChangeList {
	changeType := ChangeTypeFromString(ct)
	if changeType == Unknown {
		return nil
	}
	return &ChangeList{Type: changeType}
}

// RenderItems renders all the items
func (c *ChangeList) RenderItems(w io.Writer) {
	for _, i := range c.Items {
		i.Render(w)
	}
}

// Render builds the representation of Change
func (c *ChangeList) Render(w io.Writer) {
	io.WriteString(w, fmt.Sprintf("### %s\n", c.Type.String()))
	c.RenderItems(w)
}
