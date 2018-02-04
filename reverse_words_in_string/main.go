package main

import (
	"bufio"
	"fmt"
	"strings"
)

func reverseString(s string) {
	in := bufio.NewReader(strings.NewReader(s))
	scanner := bufio.NewScanner(in)
	//we ask scanner to split input by words for us
	scanner.Split(bufio.ScanWords)
	count := 0
	for scanner.Scan() {
		//get input token - in our case a word and update it's frequence
		word := scanner.Text()
		wordr := reverseWord(word)
		fmt.Print(wordr + " ")
		count++
	}
}

func reverseWord(s string) string {
	newstring := make([]rune, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		newstring = append(newstring, rune(s[i]))
	}
	return string(newstring)
}

func main() {
	str := "test main programm"
	fmt.Printf("Before: %s\n", str)
	reverseString(str)
}
