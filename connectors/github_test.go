package connectors

import (
	"testing"
)

func TestGithubOrganisations(t *testing.T) {

	k, err := GithubOrgs()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}

func TestGithubRepos(t *testing.T) {

	k, err := GithubRepos()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
func TestGithubCodespaces(t *testing.T) {

	k, err := GithubCodespaces()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
