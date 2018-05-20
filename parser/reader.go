package parser

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"strings"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Version stores information about the version being defined and
// its sections
type Version struct {
	Name     string
	Date     string
	Link     string
	Yanked   bool
	Sections []*Section
}

// Section holds the information about each section
// Valid sections are "Added", "Changed", "Deprecated", "Fixed",
// "Removed" and "Security"
type Section struct {
	Title   string
	Content string
}

// Reader is the implementation of blackfriday.Renderer interface
// It parses and store information about the changelog
type Reader struct {
	blackfriday.Renderer
	Versions []*Version
}

var reVersion *regexp.Regexp

func init() {
	// TODO: Make it parametrizable
	reVersion = regexp.MustCompile(`(?i)\b(v?(\d+\.?)+\b|unreleased)`)
}

// NewReader creates a new Reader instance
func NewReader(input []byte) []byte {
	extensions := blackfriday.NoIntraEmphasis

	r := Reader{}
	output := blackfriday.Run(input, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(&r))

	/*versions := ""
	for _, v := range r.Versions {
		versions += v.Name
		if v.Yanked {
			versions += "[y]"
		}
		for _, s := range v.Sections {
			versions += fmt.Sprintf("=%s=", s.Title)
			versions += fmt.Sprintf("\n%s\n", s.Content)
		}
		versions += "\n"
	}

	return []byte(fmt.Sprintf("%s\n\n%s", versions, output))*/

	return output
}

// RenderHeader is called at the beginning of the parsing
func (r *Reader) RenderHeader(w io.Writer, ast *blackfriday.Node) {}

// RenderFooter render the links to each version
func (r *Reader) RenderFooter(w io.Writer, ast *blackfriday.Node) {
	for _, v := range r.Versions {
		io.WriteString(w, fmt.Sprintf("[%s]: %s\n", v.Name, v.Link))
	}
}

// RenderNode is called for every node on the AST tree
func (r *Reader) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	switch node.Type {
	case blackfriday.Heading:
		return r.Heading(w, node, entering)
	case blackfriday.List:
		return r.List(w, node, entering)
	case blackfriday.Item:
		return r.ListItem(w, node, entering)
	case blackfriday.Code:
		return r.Code(w, node, entering)
	case blackfriday.Text:
		return r.Text(w, node, entering)
	case blackfriday.Paragraph:
		return r.Paragraph(w, node, entering)
	case blackfriday.Link:
		return r.Link(w, node, entering)
	}
	return blackfriday.GoToNext
}

// Heading is called for each Heading (h1 to h6) node found
func (r *Reader) Heading(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	level := node.HeadingData.Level
	if entering == true {
		// It's a version
		if level == 2 {
			v := Version{}
			// we append it before because Link needs it
			r.Versions = append(r.Versions, &v)

			buf := r.children(node, entering)
			line := string(buf.Bytes())
			version := reVersion.FindString(line)
			if version != "" {
				v.Name = version
				if strings.HasSuffix(line, "[YANKED]") {
					v.Yanked = true
				}
			} else {
				// now we remove it if don't needed
				r.Versions = r.Versions[:len(r.Versions)-1]
			}
			io.WriteString(w, strings.Repeat("#", level)+" ")
			io.WriteString(w, line)
			io.WriteString(w, "\n")
			return blackfriday.SkipChildren
		}
		// It's a section
		if level == 3 {
			// Get current version
			if v := r.currentVersion(); v != nil {
				buf := r.children(node, entering)
				title := string(buf.Bytes())
				v.Sections = append(v.Sections, &Section{Title: title})
			}
		}

		io.WriteString(w, strings.Repeat("#", level)+" ")
	} else {
		w.Write([]byte{'\n'})
	}
	return blackfriday.GoToNext
}

// List is called at list boundaries
func (r *Reader) List(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if entering {
		if s := r.currentSection(); s != nil {
			buf := r.children(node, entering)
			s.Content = string(buf.Bytes())
			// Uncomment when disabling output
			// return blackfriday.SkipChildren
		}
	} else {
		io.WriteString(w, "\n")
	}
	return blackfriday.GoToNext
}

// ListItem is called for each item
func (r *Reader) ListItem(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if entering {
		io.WriteString(w, "- ")
	} else {
		io.WriteString(w, "\n")
	}
	return blackfriday.GoToNext
}

// Code handles inline code marks
func (r *Reader) Code(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	w.Write([]byte{'`'})
	w.Write(node.Literal)
	w.Write([]byte{'`'})
	return blackfriday.SkipChildren
}

// Paragraph handles... paragraphs
func (r *Reader) Paragraph(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// Item will handle it's own spacing stuff
	if entering == false {
		if node.Parent.Type != blackfriday.Item {
			io.WriteString(w, "\n\n")
		}
	}

	return blackfriday.GoToNext
}

// Text renders text nodes
func (r *Reader) Text(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// Check if has \n and parent is Item
	// then add 2 spaces before the \n
	// the order is Item > Paragraph > Text
	output := node.Literal
	if node.Parent != nil && node.Parent.Parent != nil && node.Parent.Parent.Type == blackfriday.Item {
		lines := bytes.Split(node.Literal, []byte{'\n'})
		output = bytes.Join(lines, []byte("\n  "))
	}
	w.Write(output)
	return blackfriday.GoToNext
}

// Link deals with hyperlinks inside
func (r *Reader) Link(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if entering {
		io.WriteString(w, "[")
	} else {
		io.WriteString(w, "]")
		// For versions, store and print on the footer
		if node.Parent.Type == blackfriday.Heading && node.Parent.HeadingData.Level == 2 {
			if v := r.currentVersion(); v != nil {
				v.Link = string(node.LinkData.Destination)
			}
		} else {
			s := fmt.Sprintf("(%s)", node.LinkData.Destination)
			io.WriteString(w, s)
		}
	}
	return blackfriday.GoToNext
}

func (r *Reader) children(node *blackfriday.Node, entering bool) bytes.Buffer {
	var buf bytes.Buffer
	for n := node.FirstChild; n != nil; n = n.Next {
		n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
			return r.RenderNode(&buf, node, entering)
		})
	}
	return buf
}

func (r *Reader) currentVersion() *Version {
	if len(r.Versions) > 0 {
		return r.Versions[len(r.Versions)-1]
	}

	return nil
}

func (r *Reader) currentSection() *Section {
	if v := r.currentVersion(); v != nil {
		if len(v.Sections) > 0 {
			return v.Sections[len(v.Sections)-1]
		}
	}

	return nil
}
