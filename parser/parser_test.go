package parser_test

import (
	"bufio"
	"fmt"
	"os"
	"testing"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/stretchr/testify/assert"
)

func readFile(t *testing.T, filename string) *bufio.Reader {
	input, err := os.Open(fmt.Sprintf("testdata/%s.md", filename))
	if err != nil {
		t.Errorf("Failed to read testdata: %s", err)
	}

	return bufio.NewReader(input)
}

func TestParserParse(t *testing.T) {
	t.Run("simple-scenario", func(t *testing.T) {
		input := readFile(t, "simple")

		expected := &chg.Changelog{
			Preamble: "Simple paragraph.",
			Versions: []*chg.Version{
				{
					Name: "Unreleased",
					Link: "http://example.com/1.0.0..HEAD",
					Changes: []*chg.ChangeList{
						{
							Type: chg.Added,
							Items: []*chg.Item{
								{"Awesome feature that people always asked for"},
							},
						},
						{
							Type: chg.Fixed,
							Items: []*chg.Item{
								{"That annoying bug"},
							},
						},
					},
				},
				{
					Name:   "1.0.0",
					Date:   "2018-04-23",
					Link:   "http://example.com/abcdef..1.0.0",
					Yanked: true,
					Changes: []*chg.ChangeList{
						{
							Type: chg.Security,
							Items: []*chg.Item{
								{"Remote code execution using our eval endpoint"},
							},
						},
					},
				},
			},
		}

		result := parser.Parse(input)
		assert.Equal(t, expected, result)
	})

	t.Run("formatting", func(t *testing.T) {
		input := readFile(t, "formatting")
		expected := &chg.Changelog{
			Preamble: `Nesciunt **voluptate** qui _consequatur_ eos_velit quia_aut. Qui
repellendus ~~et~~ impedit ` + "`minus`" + ` inventore. Dolorem numquam voluptate
accusamus ut nihil. Aut quasi dolores quod accusamus provident facilis.
Dolores et quidem consequatur qui sequi consequatur id. Magnam ea iure
est officia [est](http://example.com).`,
		}

		result := parser.Parse(input)
		assert.Equal(t, expected, result)
	})

	t.Run("keepachangelog", func(t *testing.T) {
		input := readFile(t, "keepachangelog")

		expectedPreamble := `All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/)
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).`

		unreleasedVersion := &chg.Version{
			Name:    "Unreleased",
			Date:    "",
			Link:    "https://github.com/olivierlacan/keep-a-changelog/compare/v1.0.0...HEAD",
			Changes: nil,
		}

		zerozerooneVersion := &chg.Version{
			Name: "0.0.1",
			Date: "2014-05-31",
			Link: "",
			Changes: []*chg.ChangeList{
				{
					Type: chg.Added,
					Items: []*chg.Item{
						{Description: "This CHANGELOG file to hopefully serve as an evolving example of a\nstandardized open source project CHANGELOG."},
						{Description: "CNAME file to enable GitHub Pages custom domain"},
						{Description: "README now contains answers to common questions about CHANGELOGs"},
						{Description: "Good examples and basic guidelines, including proper date formatting."},
						{Description: "Counter-examples: \"What makes unicorns cry?\""},
					},
				},
			},
		}

		result := parser.Parse(input)
		assert.Equal(t, expectedPreamble, result.Preamble)
		assert.Equal(t, 13, len(result.Versions))
		assert.Equal(t, unreleasedVersion, result.Versions[0])
		assert.Equal(t, zerozerooneVersion, result.Versions[len(result.Versions)-1])
	})

	t.Run("malformed", func(t *testing.T) {
		input := readFile(t, "malformed")

		expected := &chg.Changelog{
			Preamble: "Simple paragraph.",
			Versions: []*chg.Version{
				{
					Name: "Unreleased",
					Link: "http://example.com/1.0.0..HEAD",
					Changes: []*chg.ChangeList{
						{
							Type: chg.Added,
							Items: []*chg.Item{
								{"Awesome feature that people always asked for"},
							},
						},
						{
							Type: chg.Fixed,
							Items: []*chg.Item{
								{"That annoying bug"},
							},
						},
					},
				},
				{
					Name: "1.0.0",
					Date: "2018-04-23",
					Link: "http://example.com/abcdef..1.0.0",
					Changes: []*chg.ChangeList{
						{
							Type: chg.Security,
							Items: []*chg.Item{
								{"Remote code execution using our eval endpoint"},
							},
						},
					},
				},
			},
		}

		result := parser.Parse(input)
		assert.Equal(t, expected, result)
	})

	t.Run("duplicated", func(t *testing.T) {
		input := readFile(t, "duplicated")

		expected := &chg.Changelog{
			Preamble: "Simple paragraph.",
			Versions: []*chg.Version{
				{
					Name: "Unreleased",
					Link: "http://example.com/abcdef..HEAD",
					Changes: []*chg.ChangeList{
						{
							Type: chg.Added,
							Items: []*chg.Item{
								{"Item 1"},
								{"Item 2"},
							},
						},
					},
				},
			},
		}

		result := parser.Parse(input)
		assert.Equal(t, expected, result)
	})
}
