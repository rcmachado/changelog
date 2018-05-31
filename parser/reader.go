package parser

import (
	"bytes"
	"fmt"
	"io"
	"regexp"

	"github.com/rcmachado/changelog/chg"

	blackfriday "gopkg.in/russross/blackfriday.v2"
)

// Reader is the implementation of blackfriday.Renderer interface
// It parses the changelog file and populate correct structs
type Reader struct {
	blackfriday.Renderer
	Changelog *chg.Changelog

	isInPreamble bool
	preambleBuf  bytes.Buffer
}

var reVersionLine *regexp.Regexp

// Points to the name of the current version being populated
var currentVersion *chg.Version

// Type of current changelist
var currentChangeType chg.ChangeType

func init() {
	// parses "version - date [yanked]" (only version is required)
	reVersionLine = regexp.MustCompile(`(?i)\[?(?P<name>[0-9a-zA-Z\-\.]+)\]?(?: - (?P<date>[0-9a-z\-\.]+))?(?P<yanked> \[YANKED\])?`)
}

// parseVersionHeading returns a filled Version object from the heading
func parseVersionHeading(line string) *chg.Version {
	submatches := reVersionLine.FindStringSubmatch(line)
	totalMatches := len(submatches)
	if totalMatches < 2 {
		return nil
	}

	v := &chg.Version{}
	v.Name = submatches[1]

	if totalMatches > 2 {
		v.Date = submatches[2]
	}
	if totalMatches > 3 && submatches[3] != "" {
		v.Yanked = true
	}

	return v
}

// Parse formats the input following the recommendation
func Parse(input []byte) *chg.Changelog {
	extensions := blackfriday.NoIntraEmphasis

	r := Reader{}
	r.Changelog = &chg.Changelog{}

	blackfriday.Run(input, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(&r))

	return r.Changelog
}

// RenderHeader is called at the beginning of the parsing
func (r *Reader) RenderHeader(w io.Writer, ast *blackfriday.Node) {}

// RenderFooter render the links to each version
func (r *Reader) RenderFooter(w io.Writer, ast *blackfriday.Node) {
	r.Changelog.RenderLinks(w)
}

// RenderNode is called for every node on the AST tree
func (r *Reader) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// Store it for future reference
	var writer io.Writer
	if r.isInPreamble {
		writer = &r.preambleBuf
	} else {
		writer = w
		// if buf has something, flush it
		content := r.preambleBuf.Bytes()
		if len(content) > 0 {
			r.Changelog.Preamble = string(content)
			r.preambleBuf.Reset()
		}
	}

	switch node.Type {
	case blackfriday.Heading:
		// overwrite buffer
		if node.Type == blackfriday.Heading && node.HeadingData.Level == 2 {
			writer = w
		}
		return r.Heading(writer, node, entering)
	case blackfriday.Item:
		return r.ListItem(writer, node, entering)
	case blackfriday.Code:
		return r.Code(writer, node, entering)
	case blackfriday.Text:
		return r.Text(writer, node, entering)
	case blackfriday.Paragraph:
		return r.Paragraph(writer, node, entering)
	case blackfriday.Link:
		return r.Link(writer, node, entering)
	}
	return blackfriday.GoToNext
}

// Heading is called for each Heading (h1 to h6) node found
func (r *Reader) Heading(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	level := node.HeadingData.Level
	if entering == true {
		if level == 1 {
			// buf := r.children(node, entering)
			// r.Changelog.Title = string(buf.Bytes())
			r.isInPreamble = true
			return blackfriday.SkipChildren
		}
		// It's a version
		if level == 2 {
			r.isInPreamble = false

			buf := r.children(node, entering)
			line := string(buf.Bytes())
			v := parseVersionHeading(line)
			r.Changelog.Versions = append(r.Changelog.Versions, v)
			if v != nil {
				currentVersion = v
				currentChangeType = 0
			} else {
				// now we remove it if don't needed
				r.Changelog.Versions = r.Changelog.Versions[:len(r.Changelog.Versions)-1]
			}
			return blackfriday.SkipChildren
		}
		// It's a section
		if level == 3 {
			// Get current version
			if currentVersion != nil {
				buf := r.children(node, entering)
				title := string(buf.Bytes())
				tmpChange := chg.NewChangeList(title)
				change := currentVersion.Change(tmpChange.Type)
				if change == nil {
					change = tmpChange
				}
				currentChangeType = change.Type
				currentVersion.Changes = append(currentVersion.Changes, change)
			}
		}
	}
	return blackfriday.GoToNext
}

// ListItem is called for each item
func (r *Reader) ListItem(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if currentVersion != nil {
		if s := currentVersion.Change(currentChangeType); s != nil {
			buf := r.children(node, true)
			s.Items = append(s.Items, &chg.Item{Description: buf.String()})
		}
	}

	return blackfriday.SkipChildren
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
			io.WriteString(w, "\n")
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
			if currentVersion != nil {
				currentVersion.Link = string(node.LinkData.Destination)
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
