package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := flag.String("url", "", "start url")
	//timeout := flag.Int("timeout", 2, "number of seconds between requests")
	flag.Parse()

	if *url == "" {
		log.Fatal("--url paramters is required")
	}

	resp, err := http.Get(*url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	fmt.Printf("%s downloaded\n", *url)
	fmt.Printf("status code: %s\n", resp.Status)
	fmt.Printf("body: %s\n", body)
}
