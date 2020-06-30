package main

import (
	"flag"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
)

type Data struct {
	Content string
	// List []entry
}

func writeFile(fileName string) *os.File {
	fileName = strings.Split(fileName, ".")[0] + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}

	return file
}

func readFromFile(file string) string {
	fileContents, err := ioutil.ReadFile(file)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	return string(fileContents)
}

func writeTemplate(file string) {
	var fileData Data
	fileData.Content = readFromFile(file)

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile := writeFile(file)

	err := t.Execute(newFile, fileData)
	if err != nil {
		panic(err)
	}
}

func main() {

}

func save() {

	fileFlag := flag.String("file", "second-post.txt", "file name you want to use for content")
	flag.Parse()

	writeTemplate(*fileFlag)
}
