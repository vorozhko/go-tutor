package helloworld

import "fmt"

// Name - define default prefix for sayName function - see in world.go
const Name = "Hello World!"

// say - print string
func say(str string) {
	fmt.Print(str)
}
