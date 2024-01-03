package connectors

import (
	"encoding/json"
	"log"
	"strings"
)

func GithubOrgs() ([]Connector, error) {
	mateContext, err := GetMateContext()
	bytes, err := Execute("gh", *&Options{}, "org", "list", "-L", "40")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	orgs := strings.Split(string(bytes), "\n")
	for _, org := range orgs {
		name := strings.TrimSpace(org)
		if name == "" {
			continue
		}
		connector := Connector{
			Name:      name,
			Url:       "https://github.com/" + name,
			IsCurrent: mateContext.Current.GitOrg == name,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}

func GithubRepos() ([]Connector, error) {
	type Repo struct {
		Name        string `json:"name"`
		ID          string `json:"id"`
		Description string `json:"description"`
		Url         string `json:"url"`
	}
	mateContext, err := GetMateContext()
	bytes, err := Execute("gh", *&Options{}, "repo", "list", mateContext.Current.GitOrg, "-L", "30", "--json", "id,name,description,url")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	repos := make([]Repo, 0)

	err = json.Unmarshal(bytes, &repos)
	if err != nil {
		return nil, err
	}
	for _, repo := range repos {
		connector := Connector{
			Name:        repo.Name,
			Url:         repo.Url,
			Description: repo.Description,
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}

func GithubCodespaces() ([]Connector, error) {
	type GitStatus struct {
		Ahead                 int64  `json:"ahead"`
		Behind                int64  `json:"behind"`
		HasUncommittedChanges bool   `json:"hasUncommittedChanges"`
		HasUnpushedChanges    bool   `json:"hasUnpushedChanges"`
		Ref                   string `json:"ref"`
	}

	type CodeSpace struct {
		Name        string    `json:"name"`
		DisplayName string    `json:"displayName"`
		GitStatus   GitStatus `json:"gitStatus"`
		Owner       string    `json:"owner"`
		Repository  string    `json:"repository"`
	}
	bytes, err := Execute("gh", *&Options{}, "codespace", "list", "-L", "30", "--json", "name,displayName,gitStatus,owner,repository")
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)

	var codespaces []CodeSpace = make([]CodeSpace, 0)

	err = json.Unmarshal(bytes, &codespaces)
	if err != nil {
		log.Fatal(string(bytes))
		return nil, err
	}
	for _, codespace := range codespaces {
		connector := Connector{
			Name: codespace.DisplayName,
			Url:  "https://" + codespace.Name + ".github.dev",
		}
		connectors = append(connectors, connector)
	}

	return connectors, nil
}
