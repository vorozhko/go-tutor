package main

//todo:
// - every link must be visited only once
// - keep a map of visited links
// - fix links and page concatination e.g. https://golang.org/doc/blog//doc/tos.html
// - extract domain from request uri to simplify crawling
// - fix how crawling settings are set like depth and maxLinks

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

//hash of visited links to prevent double visit
var visitedLinks map[string]bool
var baseURL = flag.String("url", "", "start url")

func main() {
	visitedLinks = make(map[string]bool)
	flag.Parse()

	if *baseURL == "" {
		log.Fatal("--url paramters is required")
	}

	visitedLinks[*baseURL] = false
	messages := make(chan string)
	go crawl("/", 0, 10, 10, messages)

	for {
		fmt.Print(<-messages)
	}
}

func printLinks(link string, c chan string) {
	c <- fmt.Sprintf("Crawling %s ..................\n\n", link)
}

func crawl(link string, currentDepth, finalDepth, maxLinks int, c chan string) {
	//mux.Lock()
	if visitedLinks[link] {
		return
	}
	visitedLinks[link] = true
	//mux.Unlock()
	if currentDepth == finalDepth {
		return
	}
	printLinks(link, c)
	//	fmt.Printf("Crawling %s ..................\n\n", *baseURL+link)
	resp, err := http.Get(*baseURL + link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	linkCounter := 0
	//todo: every link must be visited only once
	for _, href := range getLinks(resp.Body) {
		//crawl only internal links
		if string(href[0]) == "/" && // only internal links
			href != link { //skip current page
			if len(href) > 1 && href[1] == '/' { //skip external links which start with //
				continue
			}
			if linkCounter == maxLinks {
				return
			}
			linkCounter++
			//fmt.Printf("Found: %s\n", href)
			crawl(href, currentDepth+1, finalDepth, maxLinks, c)
		}
	}
}

//Collect all links from response body and return it as an array of strings
func getLinks(body io.Reader) []string {
	var links []string
	z := html.NewTokenizer(body)
	for {
		tt := z.Next()

		switch tt {
		case html.ErrorToken:
			//todo: links list shoudn't contain duplicates
			return links
		case html.StartTagToken, html.EndTagToken:
			token := z.Token()
			if "a" == token.Data {
				for _, attr := range token.Attr {
					if attr.Key == "href" {
						links = append(links, attr.Val)
					}

				}
			}

		}
	}
}
