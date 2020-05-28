package main

import (
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/fatih/color"
	blackfriday "github.com/russross/blackfriday/v2"
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
	mdcontent := readFile(path)
	htmlcontent := string(blackfriday.Run([]byte(mdcontent)))
	return htmlcontent, nil
}
