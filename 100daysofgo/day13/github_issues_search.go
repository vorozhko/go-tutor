package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gopl.io/ch4/github"
)

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	var dateformat string
	for _, item := range result.Items {
		days := int(time.Since(item.CreatedAt).Hours()) / 24
		if days < 30 {
			dateformat = "less than a month old"
		} else if days < 365 {
			dateformat = "less than a year old"
		} else {
			dateformat = "more than a year old"
		}
		fmt.Printf("%s, #%-5d %9.9s %.55s\n", dateformat, item.Number, item.User.Login, item.Title)
	}
}
