package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/template"
	"time"

	"github.com/fatih/color"
)

type Data struct {
	Content string
	// List []entry
}

// returns all files within a given directory
func readFilesFromDir(dir string) []os.FileInfo {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	return files
}

// check if fileName is a .txt file return a bool
func DoesFileExist(fileName string) bool {
	if strings.Contains(fileName, ".") {
		return strings.Split(fileName, ".")[1] == "txt"
	} else {
		return false
	}
}

// creates a fileName.html file using given fileName
func writeFile(fileName string) *os.File {
	fileName = strings.Split(fileName, ".")[0] + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// returns the created file
	return file
}

// read content af a given fileName
func readFromFile(fileName string) string {
	fileContents, err := ioutil.ReadFile(fileName)
	if err != nil {
		// A common use of `panic` is to abort if a function returns an error
		// value that we don’t know how to (or want to) handle. This example
		// panics if we get an unexpected error when creating a new file.
		panic(err)
	}
	return string(fileContents)
}

func writeTemplate(fileName, translate string) {
	var fileData Data
	fileData.Content = readFromFile(fileName)

	// if translate != "" {
	// 	fileData.Content = trans.TranslateText(translate, fileData.Content)
	// }

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFile := writeFile(fileName)

	// print html to stdout
	// errs := t.Execute(os.Stdout, fileData)		uncomment to print template to stdout
	// if errs != nil {								uncomment to print template to stdout
	// 	panic(errs)									uncomment to print template to stdout
	// }											uncomment to print template to stdout
	// print html to a file
	err := t.Execute(newFile, fileData)
	if err != nil {
		panic(err)
	}
}

func main() {
	save()
}

func print(pageCount int, runtime time.Duration) {
	// fmt.Print()
	green := color.New(color.FgGreen).PrintfFunc()
	green("Success! ")
	fmt.Print("Generated ")
	count := color.New(color.Bold, color.FgWhite).PrintFunc()
	count(pageCount)
	fmt.Println(" pages in", runtime)
}

func save() {

	fileFlag := flag.String("file", "second-post.txt", "file name you want to use for content")
	dirFlag := flag.String("dir", "", "Directory name that contain all your .txt files")
	translateFlag := flag.String("lang", "", "Language you want to translate the content to.")

	flag.Parse()

	pageCount := 0

	if *dirFlag != "" {
		files := readFilesFromDir(*dirFlag)

		for _, file := range files {

			fileName := file.Name()
			// check if this filename is a txt file
			if DoesFileExist(fileName) == true {
				fmt.Println(".txt file found in your dir", fileName)
				start := time.Now()
				writeTemplate(fileName, *translateFlag)
				print(pageCount, time.Since(start))
				pageCount = pageCount + 1
			}
		}

	} else {
		start := time.Now()
		writeTemplate(*fileFlag, *translateFlag)
		fmt.Println("Time function run ")
		print(pageCount, time.Since(start))

		pageCount = pageCount + 1

	}

}
