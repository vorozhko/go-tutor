package main

import (
	"fmt"
	"sort"
)

func main() {
	m := map[string]int{"A": 1, "B": 2, "AA": 1, "C": 3, "BBB": 2}

	//reverse keys and values of m map
	var mapKeys []string
	fmt.Print("Unsorted:\n")
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
		mapKeys = append(mapKeys, k)
	}

	sort.Strings(mapKeys)
	fmt.Print("Sorted:\n")
	for _, v := range mapKeys {
		fmt.Printf("%s -> %d\n", v, m[v])
	}
}
