package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

func main() {
	name := flag.String("user", "", "user name")
	flag.Parse()
	fmt.Printf("user name is %s\n", *name)

	var userArg string
	for index, arg := range os.Args {
		pattern := "-user="
		x := strings.Index(arg, pattern)
		if x > -1 {
			userArg = arg[x+len(pattern):]
			continue
		}
		if arg == "-user" || arg == "--user" {
			userArg = os.Args[index+1]
			continue
		}
	}
	fmt.Printf("user name is %s", userArg)
}
