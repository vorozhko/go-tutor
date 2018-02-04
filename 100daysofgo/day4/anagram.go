package main

import (
	"flag"
	"fmt"
)

func main() {
	str1 := flag.String("first", "", "first string for anagram check")
	str2 := flag.String("second", "", "second string for anagram check")
	flag.Parse()
	checkForAnagrams(*str1, *str2)
}

func checkForAnagrams(str1, str2 string) {
	if len(str1) != len(str2) {
		fmt.Printf("%s and %s are not anagrams", str1, str2)
		return
	}

	for i := 0; i < len(str1); i++ {
		if str1[i] != str2[len(str2)-i-1] {
			fmt.Printf("%s and %s are not anagrams", str1, str2)
			return
		}
	}
	fmt.Printf("%s and %s are anagrams", str1, str2)
}
