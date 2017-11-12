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

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		data, err := ioutil.ReadFile(file.Name())
		fileType := http.DetectContentType(data)
		if err != nil {
			log.Fatal(err)
		}
		for _, line := range strings.Split(string(data), "\n") {
			if strings.Index(line, *search) > -1 {
				if strings.Index(fileType, "text/plain") > -1 {
					fmt.Printf("%s: %s\n", file.Name(), line)
				} else {
					fmt.Printf("Binary file %s matches\n", file.Name())
					break
				}

			}
		}
	}
}
