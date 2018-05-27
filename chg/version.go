package chg

import (
	"io"
	"sort"
)

// Version stores information about the version being defined and
// its sections
type Version struct {
	Name    string
	Date    string // Date in the format YYYY-MM-DD
	Link    string
	Yanked  bool // True if the release was yanked/removed
	Changes []*Change
}

// Change returns the Change with name
func (v *Version) Change(ct ChangeType) *Change {
	for _, c := range v.Changes {
		if c.Type == ct {
			return c
		}
	}

	return nil
}

// SortChanges sort the changes ascending
func (v *Version) SortChanges() {
	sort.Slice(v.Changes, func(i, j int) bool {
		return v.Changes[i].Type < v.Changes[j].Type
	})
}

// RenderTitle writes the title in correct format
func (v *Version) RenderTitle(w io.Writer) {
	io.WriteString(w, "## ")
	if v.Link != "" {
		io.WriteString(w, "[")
		io.WriteString(w, v.Name)
		io.WriteString(w, "]")
	} else {
		io.WriteString(w, v.Name)
	}
	if v.Date != "" {
		io.WriteString(w, " - ")
		io.WriteString(w, v.Date)
	}
	if v.Yanked {
		io.WriteString(w, " [YANKED]")
	}
}

// RenderChanges writes all the changes
func (v *Version) RenderChanges(w io.Writer) {
	for i, c := range v.Changes {
		if i > 0 {
			io.WriteString(w, "\n")
		}
		c.Render(w)
	}
}

// Render writes the title and changes
func (v *Version) Render(w io.Writer) {
	v.RenderTitle(w)
	io.WriteString(w, "\n")
	v.RenderChanges(w)
}
