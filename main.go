package main

import "github.com/rcmachado/changelog/cmd"

func main() {
	cmd.Execute()
}

/*func main() {
	input, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	extensions := blackfriday.NoIntraEmphasis | blackfriday.NoEmptyLineBeforeBlock | blackfriday.Footnotes

	renderer := &keepachangelog.Renderer{}

	output := blackfriday.Run(input, blackfriday.WithExtensions(extensions), blackfriday.WithRenderer(renderer))
	fmt.Printf("%s", output)
}*/
