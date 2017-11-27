package main

import (
	"fmt"
	"sort"
)

func main() {
	var m = make(map[string]int)
	m["Sun"] = 70
	m["Sat"] = 6
	m["Fri"] = 5
	m["Thu"] = 4
	m["Mon"] = 1
	m["Tue"] = 2
	m["Wed"] = 3

	//reverse keys and values of m map
	var days = make(map[int]string)
	//array of indexes for sorting
	var daykeys = make([]int, len(m))

	fmt.Print("Unsorted days of week\n")
	counter := 0
	for k, v := range m {
		fmt.Printf("%s -> %d\n", k, v)
		days[v] = k
		daykeys[counter] = v
		counter++
	}
	//sort indexes array here
	sort.Ints(daykeys)

	fmt.Print("Sorted days of week\n")
	for _, v := range daykeys {
		fmt.Printf("%s\n", days[v])
	}
}
