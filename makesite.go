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
}

type Functime struct {
	name      string
	func_time float64
	pageCount int
	func_size float64
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
func writeFile(fileName string) (string, *os.File) {
	fileName = strings.Split(fileName, ".")[0] + ".html"
	file, err := os.Create(fileName)
	if err != nil {
		panic(err)
	}
	// returns the created file
	return fileName, file
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

func writeTemplate(fileName, translate string) string {
	var fileData Data
	fileData.Content = readFromFile(fileName)

	t := template.Must(template.New("template.tmpl").ParseFiles("template.tmpl"))
	newFileName, newFile := writeFile(fileName)

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

	return newFileName

}

func main() {
	save()
}

func templateRecord(templateName string, runtime time.Duration) Functime {
	var templateInfo Functime
	templateInfo.name = templateName
	templateInfo.func_time = float64(runtime) / float64(time.Millisecond)
	templateInfo.pageCount = templateInfo.pageCount + 1
	// fi, err := templateFile.Stat()
	fi, err := os.Stat(templateName)
	if err != nil {
		// Could not obtain stat, handle error
	}
	templateInfo.func_size = float64(fi.Size())

	// fmt.Printf("The file is %d bytes long",)

	return templateInfo
}

func save() {

	fileFlag := flag.String("file", "second-post.txt", "file name you want to use for content")
	dirFlag := flag.String("dir", "", "Directory name that contain all your .txt files")
	translateFlag := flag.String("lang", "", "Language you want to translate the content to.")

	flag.Parse()

	var fTCount []Functime

	if *dirFlag != "" {
		files := readFilesFromDir(*dirFlag)

		for _, file := range files {

			fileName := file.Name()
			// check if this filename is a txt file
			if DoesFileExist(fileName) == true {
				// fmt.Println(".txt file found in your dir", fileName)
				start := time.Now()
				templateName := writeTemplate(fileName, *translateFlag)
				fTCount = append(fTCount, templateRecord(templateName, time.Since(start)))

			}
		}

	} else {
		start := time.Now()
		templateName := writeTemplate(*fileFlag, *translateFlag)
		fTCount = append(fTCount, templateRecord(templateName, time.Since(start)))

	}
	green := color.New(color.FgGreen).PrintfFunc()
	cyan := color.New(color.FgCyan).PrintfFunc()
	red := color.New(color.FgRed).PrintfFunc()
	boldFont := color.New(color.Bold, color.FgWhite).PrintFunc()

	green("Success! ")
	fmt.Print("Generated ")
	boldFont(len(fTCount))

	fmt.Println(" templates in: ")
	for _, record := range fTCount {
		// fmt.Println("_________ _ __ __________")
		cyan("Template: ")
		boldFont(record.name)
		fmt.Print(" ")
		boldFont(record.func_size)
		boldFont("kb")
		red(" in ")
		fmt.Printf("%.2f", record.func_time)
		// fmt.Println("_________|_|__|_________|")
		fmt.Println("ms")

	}

}
