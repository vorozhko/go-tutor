package main

//todo:
// - every link must be visited only once [done]
// - keep a map of visited links [done]
// - fix links and page concatination [done]
// - extract domain from request uri to simplify crawling [done]
// - rework how internal links are selected for crawling
// - fix how crawling settings are set like depth and maxLinks

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

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

	//set parameters for crawling
	crawl("/")
}

func crawl(link string) {
	//check if link already visited
	if visitedLinks[link] {
		return
	}
	//set link as visited
	visitedLinks[link] = true
	fmt.Printf("Crawling %s ..................\n\n", *baseURL+link)
	resp, err := http.Get(*baseURL + link)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	linkCounter := 0
	for _, href := range getLinks(resp.Body) {
		//todo: rework how links are selected
		if len(href) > 0 && string(href[0]) == "/" && // only internal links
			href != link { //skip current page
			if len(href) > 1 && href[1] == '/' { //skip external links which start with //
				continue
			}
			linkCounter++
			//fmt.Printf("Found: %s\n", href)
			crawl(href)
			time.Sleep(time.Second * 1)
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
