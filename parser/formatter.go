package parser

import blackfriday "gopkg.in/russross/blackfriday.v2"

// Formatter formats the input following the recommendation
func Formatter(input []byte) []byte {
	extensions := blackfriday.NoIntraEmphasis

	chg := Changelog{}
	r := Reader{Changelog: &chg}

	return blackfriday.Run(input, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(&r))

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
}
