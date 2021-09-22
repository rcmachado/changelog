package chg

//go:generate stringer -type=ChangeType

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"
)

// ChangeList groups the changes by type
// Valid change types are "Added", "Changed", "Deprecated", "Fixed",
// "Removed" and "Security"
type ChangeList struct {
	Type  ChangeType
	Items []*Item
}

func (cl *ChangeList) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type  string  `json:"type"`
		Items []*Item `json:"items"`
	}{
		Type:  ChangeStringFromType(cl.Type),
		Items: cl.Items,
	})
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
func ChangeStringFromType(ct ChangeType) string {
	switch ct {
	case Added:
		return "added"
	case Changed:
		return "changed"
	case Deprecated:
		return "deprecated"
	case Fixed:
		return "fixed"
	case Removed:
		return "removed"
	case Security:
		return "security"
	default:
		return "unknown"
	}
}

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
