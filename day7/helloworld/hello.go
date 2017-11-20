// Package helloworld - example of how to make a package
package helloworld

import "fmt"

// prefix - define default prefix
const prefix = "Hello World!"

// say - print string
func say(str string) {
	fmt.Print(str)
}
