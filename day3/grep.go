package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	path := flag.String("path", "", "file path to search in")
	search := flag.String("search", "", "search string to look for")
	flag.Parse()

	readFiles(*path, *search)
}

func readFiles(path, search string) {
	files, err := ioutil.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		fullpath := path + file.Name()
		if file.Mode().IsDir() {
			readFiles(fullpath+"/", search)
		} else if file.Mode().IsRegular() {
			searchInFile(fullpath, search)
		}
	}
}

func searchInFile(fullpath, search string) {
	data, err := ioutil.ReadFile(fullpath)
	if err != nil {
		log.Fatal(err)
	}
	//need to check for file type to detect binary content
	fileType := http.DetectContentType(data)
	if strings.Index(fileType, "text") == -1 {
		fmt.Printf("Not a text file %s is skiped\n", fullpath)
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Index(line, search) > -1 {
			fmt.Printf("%s: %s\n", fullpath, line)
		}
	}
}
