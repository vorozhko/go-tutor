package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	repo := flag.String("repo", "", "github owner/repo e.g. golang/go")
	id := flag.Int("id", -1, "issue id")
	flag.Parse()

	if *id == -1 || *repo == "" {
		log.Fatal("--repo and --id parameters must be provided")
	}
	issue, _ := read(*repo, *id)
	fmt.Print(issue.Title)
}

// IssueData - specify data fields for new github issue submission
type IssueData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func read(ownerRepo string, id int) (*IssueData, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d", ownerRepo, id)
	resp, err := http.Get(apiURL)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		log.Fatal(err)
		return nil, err
	}

	var result *IssueData
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}

	return result, nil
}
