package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//map to store words frequency
	words := make(map[string]int)
	//we will read from file
	in := bufio.NewReader(os.Stdin)
	scanner := bufio.NewScanner(in)
	//we ask scanner to split input by words for us
	scanner.Split(bufio.ScanWords)
	count := 0
	//scan the inpurt
	for scanner.Scan() {
		//get input token - in our case a word and update it's frequence
		words[scanner.Text()]++
		count++
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading input:", err)
	}
	fmt.Printf("Total words: %d\n", count)
	fmt.Printf("Words frequency: \n")
	//todo: sort words by values for nice print
	for k, v := range words {
		fmt.Printf("%s:%d\n", k, v)
	}
}
