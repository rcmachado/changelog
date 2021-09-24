package parser

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"regexp"
	"strings"

	"github.com/cucumber/changelog/chg"
	blackfriday "github.com/russross/blackfriday/v2"
)

// Parse input into a proper Changelog struct
func Parse(r io.Reader) *chg.Changelog {
	extensions := blackfriday.NoIntraEmphasis | blackfriday.Strikethrough
	renderer := newRenderer()

	input, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	blackfriday.Run(input, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(&renderer))

	return renderer.Result()
}

func newRenderer() renderer {
	r := renderer{}
	r.changelog = chg.NewChangelog()
	r.reVersion = regexp.MustCompile(`(?i)\[?(?P<name>[0-9a-zA-Z\-\.]+)\]?(?: - (?P<date>[0-9a-z\-\.]+))?(?P<yanked> \[YANKED\])?`)
	return r
}

type renderer struct {
	blackfriday.Renderer

	changelog      *chg.Changelog
	captureContent bool            // should we capture content in a special buffer?
	captureBuffer  bytes.Buffer    // buffer where we will store content temporarily
	reVersion      *regexp.Regexp  // matches the version line
	currentVersion *chg.Version    // current version being parsed
	currentChange  *chg.ChangeList // current changelist being parsed
}

func (r *renderer) Result() *chg.Changelog {
	return r.changelog
}

// RenderHeader is called at the beginning of the parsing
func (r *renderer) RenderHeader(w io.Writer, ast *blackfriday.Node) {}

// RenderHeader is called at the end of the parsing
func (r *renderer) RenderFooter(w io.Writer, ast *blackfriday.Node) {
	// if the changelog only has the first paragraph (without versions)
	// flush it to the
	if content := r.endCapture(); content != "" {
		r.changelog.Preamble = content
	}
}

// RenderNode is called for every node on the AST tree
func (r *renderer) RenderNode(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	var writer io.Writer
	if r.captureContent {
		writer = &r.captureBuffer
	} else {
		writer = w
	}

	switch node.Type {
	case blackfriday.Code:
		return r.Code(writer, node, entering)
	case blackfriday.Del:
		return r.Del(writer, node, entering)
	case blackfriday.Emph:
		return r.Emph(writer, node, entering)
	case blackfriday.Heading:
		return r.Heading(writer, node, entering)
	case blackfriday.Item:
		return r.ListItem(writer, node, entering)
	case blackfriday.Link:
		return r.Link(writer, node, entering)
	case blackfriday.Paragraph:
		return r.Paragraph(writer, node, entering)
	case blackfriday.Strong:
		return r.Strong(writer, node, entering)
	case blackfriday.Text:
		return r.Text(writer, node, entering)
	}
	return blackfriday.GoToNext
}

// Code handles inline code marks
func (r *renderer) Code(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	w.Write([]byte{'`'})
	w.Write(node.Literal)
	w.Write([]byte{'`'})
	return blackfriday.SkipChildren
}

// Del renders strikethrough marks
func (r *renderer) Del(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	io.WriteString(w, "~~")
	return blackfriday.GoToNext
}

// Emph renders emphasis marks
func (r *renderer) Emph(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	io.WriteString(w, "_")
	return blackfriday.GoToNext
}

// Heading is called for each Heading (1 to 6) node found
func (r *renderer) Heading(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	level := node.HeadingData.Level
	switch level {
	case 1: // Document title
		r.startCapture()
		// We don't care about changelog title
		return blackfriday.SkipChildren
	case 2: // It's a version
		if content := r.endCapture(); content != "" {
			r.changelog.Preamble = content
		}

		r.currentChange = nil
		r.currentVersion = &chg.Version{}

		var buf bytes.Buffer
		r.renderInline(&buf, node, entering)
		metadata := r.parseVersionLine(buf.String())

		r.currentVersion.Name = metadata["name"]
		r.currentVersion.Date = metadata["date"]
		if metadata["yanked"] != "" {
			r.currentVersion.Yanked = true
		}
		r.changelog.Versions = append(r.changelog.Versions, r.currentVersion)

		return blackfriday.SkipChildren
	case 3, 4: // It's a change
		var buf bytes.Buffer
		r.renderInline(&buf, node, entering)
		changeName := buf.String()

		changeType := chg.ChangeTypeFromString(changeName)
		if changeType != chg.Unknown {
			r.currentChange = r.currentVersion.Change(changeType)
			if r.currentChange == nil {
				r.currentChange = chg.NewChangeList(changeName)
				r.currentVersion.Changes = append(r.currentVersion.Changes, r.currentChange)
			}
		} else {
			r.currentChange = nil
		}

		return blackfriday.SkipChildren
	}
	return blackfriday.GoToNext
}

// Link deals with hyperlinks (both versions and in text)
func (r *renderer) Link(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if entering {
		io.WriteString(w, "[")
	} else {
		io.WriteString(w, "]")
		// For versions, store it
		if node.Parent.Type == blackfriday.Heading && node.Parent.HeadingData.Level == 2 && r.currentVersion != nil {
			r.currentVersion.Link = string(node.LinkData.Destination)
		} else {
			s := fmt.Sprintf("(%s)", node.LinkData.Destination)
			io.WriteString(w, s)
		}
	}
	return blackfriday.GoToNext
}

// ListItem is called for each item
func (r *renderer) ListItem(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	if r.currentChange != nil {
		var buf bytes.Buffer
		r.renderInline(&buf, node, entering)
		item := &chg.Item{Description: buf.String()}
		r.currentChange.Items = append(r.currentChange.Items, item)
	}

	return blackfriday.SkipChildren
}

// Paragraph handles... paragraphs
func (r *renderer) Paragraph(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	// Item will handle it's own spacing stuff
	if entering == false {
		if node.Parent.Type != blackfriday.Item {
			io.WriteString(w, "\n\n")
		}
	}

	return blackfriday.GoToNext
}

// Strong renders strong marks
func (r *renderer) Strong(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	io.WriteString(w, "**")
	return blackfriday.GoToNext
}

// Text renders text nodes
func (r *renderer) Text(w io.Writer, node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
	w.Write(node.Literal)
	return blackfriday.GoToNext
}

func (r *renderer) startCapture() {
	r.captureContent = true
}

func (r *renderer) endCapture() string {
	if r.captureContent {
		r.captureContent = false
		return strings.TrimSpace(r.captureBuffer.String())
	}

	return ""
}

func (r *renderer) parseVersionLine(line string) map[string]string {
	matches := r.reVersion.FindStringSubmatch(line)
	groupNames := r.reVersion.SubexpNames()

	mappedMatches := make(map[string]string)
	for idx, name := range groupNames {
		value := matches[idx]
		mappedMatches[name] = value
	}

	return mappedMatches
}

// renderInline renders the node right away
func (r *renderer) renderInline(w io.Writer, node *blackfriday.Node, entering bool) {
	for n := node.FirstChild; n != nil; n = n.Next {
		n.Walk(func(node *blackfriday.Node, entering bool) blackfriday.WalkStatus {
			return r.RenderNode(w, node, entering)
		})
	}
}
