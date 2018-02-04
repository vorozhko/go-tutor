package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"sort"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var newdata string
	for i := 0; i < 10000; i++ {
		newdata = randStringRunes(10)
		writeLineToFileSorted(newdata)
		fmt.Println(i)
	}

}

func writeLineToFileSorted(newdata string) {
	indexData, err := ioutil.ReadFile("db.txt")
	if err != nil {
		log.Fatal(err)
	}
	var newIndexData []string
	for _, line := range bytes.Split(indexData, []byte("\n")) {
		//fmt.Printf("%s\n", line)
		newIndexData = append(newIndexData, string(line))
	}
	newIndexData = append(newIndexData, string(newdata))
	sort.Strings(newIndexData)

	bytesData := []byte(strings.Join(newIndexData, "\n"))
	ioutil.WriteFile("db.txt", bytesData, 0644)
}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func randStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
