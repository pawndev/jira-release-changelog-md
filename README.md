# Jira release changelog to markdown

## what you need

- The jira version id will be in the url of your release `atlassian.net/projects/FP/versions/<the-id-here>>/tab/release-report-all-issues`
- Your email
- The password, follow this tuto: https://docs.searchunify.com/Content/Content-Sources/Atlassian-Jira-Confluence-Authentication-Create-API-Token.htm

## Launch it

After installing the go dependencies :

```shell
JIRA_VERSION_ID=15025 \
JIRA_USERNAME=jira@email.com \
JIRA_PASSWORD=<My-jira-key> go run ./cmd/main.go
```

