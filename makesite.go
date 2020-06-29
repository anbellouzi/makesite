package main

import (
	"io/ioutil"
	"os"
	"text/template"
)

type Data struct {
	Content string
	// List []entry
}

func writeFile() {
	bytesToWrite := []byte("hello\ngo\n")
	err := ioutil.WriteFile("new-file.txt", bytesToWrite, 0644)
	if err != nil {
		panic(err)
	}
}

func readFromFile() string {
	fileContents, err := ioutil.ReadFile("first-post.txt")
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	// fmt.Print(string(fileContents))

	return string(fileContents)
}

func writeTemplate() {
	var fileData Data
	fileData.Content = readFromFile()

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	err := t.Execute(os.Stdout, fileData)
	if err != nil {
		panic(err)
	}
}

func main() {

	writeTemplate()
}
