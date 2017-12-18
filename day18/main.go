package main

import (
	"crypto/md5"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("expected filename as parameter")
	}
	filename := os.Args[1]

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("MD5 (%s) = %x\n", filename, md5.Sum(data))
}
