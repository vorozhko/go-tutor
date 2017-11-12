package main

import (
	"flag"
	"fmt"
	"log"
)

func main() {
	name := flag.String("user", "", "your name")
	flag.Parse()
	if *name == "" {
		log.Fatal("--user flag is required")
	}
	fmt.Printf("Your name is %s", *name)
}
