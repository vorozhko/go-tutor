package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
)

func main() {
	token := flag.String("token", "", "github auth token")
	repo := flag.String("repo", "", "github owner/repo e.g. golang/go")
	id := flag.Int("id", -1, "issue id")
	flag.Parse()

	if *token == "" || *id == -1 || *repo == "" {
		log.Fatal("--token, --repo and --id parameters must be provided")
	}
	issue, _ := read(*repo, *id, *token)
	fmt.Print(issue.Title)
}

// IssueData - specify data fields for new github issue submission
type IssueData struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func read(ownerRepo string, id int, token string) (*IssueData, error) {
	apiURL := fmt.Sprintf("https://api.github.com/repos/%s/issues/%d", ownerRepo, id)
	fmt.Println(apiURL)
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
