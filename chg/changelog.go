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

// Release transforms Unreleased into the version informed
func (c *Changelog) Release(newVersion Version) (*Version, error) {
	oldUnreleased := c.Version("Unreleased")
	prevVersion := c.Versions[len(c.Versions)-1]

	newUnreleased := Version{
		Name: "Unreleased",
	}

	if newVersion.Link != "" {
		newUnreleased.Link = strings.Replace(oldUnreleased.Link, prevVersion.Name, newVersion.Name, -1)
	} else {
		if prevVersion == oldUnreleased {
			// we don't have a previous version
			return nil, fmt.Errorf("Could not infer the compare link")
		}
	}

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
