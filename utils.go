package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/parser"
)

func check(err error) {
	if err != nil {
		color.Red(err.Error())
	}
}

func readDir(dirpath string) []string {
	var elems []string
	_ = filepath.Walk(dirpath, func(path string, f os.FileInfo, err error) error {
		elems = append(elems, path)
		return nil
	})
	return elems
}

func readFile(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		color.Red(path)
	}
	check(err)
	return string(dat)
}

func fileToHTML(path string) (string, error) {
	mdExtensions := parser.NoIntraEmphasis | parser.Tables | parser.FencedCode |
		parser.Autolink | parser.Strikethrough | parser.SpaceHeadings | parser.HeadingIDs |
		parser.BackslashLineBreak | parser.DefinitionLists | parser.MathJax

	mdcontent := readFile(path)
	mdParser := parser.NewWithExtensions(mdExtensions)
	htmlcontent := markdown.ToHTML([]byte(mdcontent), mdParser, nil)

	// htmlcontent := string(blackfriday.Run([]byte(mdcontent)))
	return string(htmlcontent), nil
}
