package kitchen

import (
	"encoding/json"
	"log"
	"testing"
)

func TestRun(t *testing.T) {

	k, err := List()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)
	//text := string("dafdsf")

	j, err := json.Marshal(k)
	log.Println(string(j))

}

func TestReadMarkdownFromPowerShell(t *testing.T) {
	//text := string("dafdsf")

	text, err := ReadMarkdownFromPowerShell("/Users/nielsgregersjohansen/kitchens/sharepoint-branding/install/20 apply-sitetemplate.ps1")
	if err != nil {
		t.Error(err)
	}
	log.Println(text)

}

func TestGetStations(t *testing.T) {

	k, err := GetStations("sharepoint-branding")
	if err != nil {
		t.Error(err)
	}
	t.Log(k)

}
