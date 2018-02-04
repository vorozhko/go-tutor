package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

var path = flag.String("path", "", "file path to search in")
var search = flag.String("search", "", "search string to look for")

func main() {
	flag.Parse()
	fi, err := os.Stat(*path)
	if err != nil {
		log.Fatal(err)
	}
	//fix path if directory
	if fi.Mode().IsDir() {
		*path = strings.TrimRight(*path, "/") + "/"
		readFiles(*path, *search)
	} else {
		log.Fatal("path must be a directory, but file was provided: ", *path)
	}
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
	//need to check for file type to detect filter off non-text files
	fileType := http.DetectContentType(data)
	if strings.Index(fileType, "text") == -1 {
		//skip all non text files
		return
	}
	for _, line := range strings.Split(string(data), "\n") {
		if strings.Index(line, search) > -1 {
			fmt.Printf("%s: %s\n", fullpath, line)
		}
	}
}
