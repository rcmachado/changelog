package chg

import (
	"bytes"
	"fmt"
	"io"
	"strings"
)

// Changelog is the main struct that holds all the data
// in a format specific to the spec
type Changelog struct {
	Preamble string
	Versions []*Version
}

// NewChangelog creates the changelog struct
func NewChangelog() *Changelog {
	c := Changelog{}
	return &c
}

// Version finds and returns the version `v`
// The search is case-insensitive
func (c *Changelog) Version(version string) *Version {
	for _, v := range c.Versions {
		if strings.ToLower(v.Name) == strings.ToLower(version) {
			return v
		}
	}

	return nil
}

// AddItem includes the message under the proper section of Unreleased version
func (c *Changelog) AddItem(section ChangeType, message string) {
	v := c.Version("Unreleased")
	if v == nil {
		v = &Version{Name: "Unreleased"}
		c.Versions = append([]*Version{v}, c.Versions...)
	}

	s := v.Change(section)
	if s == nil {
		s = NewChangeList(section.String())
		v.Changes = append(v.Changes, s)
	}
	item := &Item{
		Description: message,
	}
	s.Items = append(s.Items, item)
}

// Release transforms Unreleased into the version informed
func (c *Changelog) Release(newVersion Version) (*Version, error) {
	oldUnreleased := c.Version("Unreleased")
	prevVersion := c.Versions[1]

	newUnreleased := Version{
		Name: "Unreleased",
	}

	if prevVersion == oldUnreleased && newVersion.Link == "" {
		// we don't have a previous version
		return nil, fmt.Errorf("Could not infer the compare link")
	}

	var compareURL string
	if newVersion.Link != "" {
		compareURL = strings.Replace(newVersion.Link, "<prev>", newVersion.Name, -1)
		compareURL = strings.Replace(compareURL, "<next>", "HEAD", -1)
	} else {
		compareURL = strings.Replace(oldUnreleased.Link, prevVersion.Name, newVersion.Name, -1)
	}

	newUnreleased.Link = compareURL

	oldUnreleased.Link = strings.Replace(oldUnreleased.Link, "HEAD", newVersion.Name, -1)
	oldUnreleased.Name = newVersion.Name
	oldUnreleased.Date = newVersion.Date

	c.Versions = append([]*Version{&newUnreleased}, c.Versions...)

	return oldUnreleased, nil
}

// RenderLinks will render the links for each version
func (c *Changelog) RenderLinks(w io.Writer) {
	for _, v := range c.Versions {
		if v.Link != "" {
			io.WriteString(w, fmt.Sprintf("[%s]: %s\n", v.Name, v.Link))
		}
	}
}

// Render outputs the full changelog contents
func (c *Changelog) Render(w io.Writer) {
	io.WriteString(w, "# Changelog\n")
	if preamble := strings.TrimSpace(c.Preamble); preamble != "" {
		io.WriteString(w, "\n")
		io.WriteString(w, preamble)
		io.WriteString(w, "\n")
	}
	for _, v := range c.Versions {
		io.WriteString(w, "\n")
		v.SortChanges()
		v.Render(w)
	}

	var buf bytes.Buffer
	c.RenderLinks(&buf)
	if content := buf.Bytes(); content != nil {
		io.WriteString(w, "\n")
		w.Write(content)
	}
}
