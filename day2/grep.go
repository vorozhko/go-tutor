package main

import (
	"flag"
	"fmt"
)

func main() {
	filepath := flag.String("filepath", "", "file path to search in")
	flag.Parse()
	fmt.Printf("Search in... %s\n", *filepath)
}
