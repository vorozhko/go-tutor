package main

import (
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	token := flag.String("token", "", "github auth token")
	repo := flag.String("repo", "", "github owner/repo e.g. golang/go")
	title := flag.String("title", "", "title for new issue")
	body := flag.String("body", "", "body for new issue")
	flag.Parse()

	if *token == "" || *title == "" || *repo == "" {
		log.Fatal("--token, --repo and --title parameters must be provided")
	}
	create(*repo, *title, *body, *token)
}

// NewIssue - specify data fields for new github issue submission
type NewIssue struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func create(ownerRepo, title, body, token string) {
	apiURL := "https://api.github.com/repos/" + ownerRepo + "/issues"
	//title is the only required field
	issueData := NewIssue{Title: title, Body: body}
	//make it json
	jsonData, _ := json.Marshal(issueData)
	//creating client to set custom headers for Authorization
	client := &http.Client{}
	req, _ := http.NewRequest("POST", apiURL, bytes.NewReader(jsonData))
	req.Header.Set("Authorization", "token "+token)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		fmt.Printf("Response code is is %d\n", resp.StatusCode)
		body, _ := ioutil.ReadAll(resp.Body)
		//print body as it may contain hints in case of errors
		fmt.Println(string(body))
		log.Fatal(err)
	}
}
