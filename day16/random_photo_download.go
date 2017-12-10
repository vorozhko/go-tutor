package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	url := "https://picsum.photos/200/300/?random"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	image, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = ioutil.WriteFile("./myimage.gif", image, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
