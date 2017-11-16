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
	"strings"

	"golang.org/x/net/html"
)

func main() {
	url := flag.String("url", "", "start url")
	//timeout := flag.Int("timeout", 2, "number of seconds between requests")
	flag.Parse()

	if *url == "" {
		log.Fatal("--url paramters is required")
	}

	crawl(*url, 0, 3, 5)
}

func crawl(url string, currentDepth, finalDepth, maxLinks int) {
	if currentDepth == finalDepth {
		return
	}
	fmt.Printf("Crawling %s ..................\n\n", url)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	linkCounter := 0
	//todo: every link must be visited only once
	for _, link := range getLinks(resp.Body) {
		//crawl only internal links
		if string(link[0]) == "/" && // only internal links
			strings.Index(url, link) == -1 { //skip current page
			if linkCounter == maxLinks {
				return
			}
			linkCounter++
			fmt.Printf("Found: %s\n", link)
			crawl(url+link, currentDepth+1, finalDepth, maxLinks)
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
