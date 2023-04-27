package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
	"os"
)

const (
	JIRA_VERSION_ID = "JIRA_VERSION_ID"
	JIRA_USERNAME   = "JIRA_USERNAME"
	JIRA_PASSWORD   = "JIRA_PASSWORD"
	JIRA_DOMAIN     = "JIRA_DOMAIN"
)

func main() {
	releaseID := os.Getenv(JIRA_VERSION_ID)

	tp := jira.BasicAuthTransport{
		Username: os.Getenv(JIRA_USERNAME),
		Password: os.Getenv(JIRA_PASSWORD),
	}
	jiraClient, _ := jira.NewClient(tp.Client(), os.Getenv(JIRA_DOMAIN))
	issues, _, err := jiraClient.Issue.Search(fmt.Sprintf("fixVersion=%s", releaseID), nil)
	if err != nil {
		panic(err)
	}

	for _, issue := range issues {
		fmt.Printf("- [%s](https://e-tf1.atlassian.net/browse/%s) %s\n", issue.Key, issue.Key, issue.Fields.Summary)
	}
}
