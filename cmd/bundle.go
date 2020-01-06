package cmd

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/rcmachado/changelog/chg"
	"github.com/rcmachado/changelog/parser"
	"github.com/spf13/cobra"
)

var directory string

var bundleCmd = &cobra.Command{
	Use:   "bundle",
	Short: "Bundles files containing unrelased changelog entries",
	Long:  "Bundles multiple files that follows the changetype/file.md structure into the Unreleased version",
	Run: func(cmd *cobra.Command, args []string) {
		changelog := parser.Parse(inputFile)
		bundleFiles(directory, changelog)
		changelog.Render(outputFile)
	},
}

func init() {
	bundleCmd.Flags().StringVarP(&directory, "directory", "d", "changelog-unreleased", "Directory where we store entries to be bundled")
	rootCmd.AddCommand(bundleCmd)
}

func bundleFiles(root string, c *chg.Changelog) {
	validTypes := map[string]chg.ChangeType{
		"added":      chg.Added,
		"changed":    chg.Changed,
		"deprecated": chg.Deprecated,
		"fixed":      chg.Fixed,
		"removed":    chg.Removed,
		"security":   chg.Security,
	}

	var currentChangeType chg.ChangeType

	filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Printf("prevent panic by handling failure accessing a path %q: %v\n", path, err)
			return err
		}

		name := strings.ToLower(info.Name())
		changeType, ok := validTypes[name]
		if info.IsDir() {
			if name != root && ok == false {
				return filepath.SkipDir
			}
			currentChangeType = changeType

		} else {
			ext := filepath.Ext(info.Name())
			if ext == ".md" {
				readFile(path, currentChangeType, c)
			}
		}
		return nil
	})
}

func readFile(path string, ct chg.ChangeType, c *chg.Changelog) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		line := strings.TrimLeft(scanner.Text(), "- ")
		c.AddItem(ct, line)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
