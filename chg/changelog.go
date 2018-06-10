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
		v.Render(w)
	}

	var buf bytes.Buffer
	c.RenderLinks(&buf)
	if content := buf.Bytes(); content != nil {
		io.WriteString(w, "\n")
		w.Write(content)
	}
}
