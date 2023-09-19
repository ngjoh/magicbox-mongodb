package powershell

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"
)

type HubSite struct {
	ID          string      `json:"ID"`
	Description interface{} `json:"Description"`
	Title       string      `json:"Title"`
	SiteURL     string      `json:"SiteUrl"`
}

type HubSites []HubSite
type Pages []struct {
	HubSiteID   string    `json:"HubSiteId"`
	Editor      string    `json:"Editor"`
	RelativeURL string    `json:"RelativeURL"`
	ID          int       `json:"ID"`
	CreatedOn   time.Time `json:"CreatedOn"`
	PageName    string    `json:"PageName"`
	ModifiedOn  time.Time `json:"ModifiedOn"`
}

type SitePages []struct {
	Title       string `json:"title"`
	Pages       Pages  `json:"pages"`
	WelcomePage string `json:"WelcomePage"`
	Siteurl     string `json:"siteurl"`
	HubSiteID   string `json:"HubSiteId"`
}
type NavigationNode struct {
	Childs      []NavigationNode `json:"Childs"`
	Title       string           `json:"Title"`
	RelativeURL string           `json:"RelativeURL"`
}
type Navigation []NavigationNode

type SiteNavigation struct {
	Navigation Navigation `json:"navigation"`
	Title      string     `json:"title"`
	Siteurl    string     `json:"siteurl"`
}
type CopyPageResult struct {
	NewPageURL string `json:"NewPageURL"`
}

func GetHubSpokesSitePages(hubId string) (*SitePages, error) {
	powershellScript := "scripts/sharepoint/get-hubsite-spokes-pages.ps1"
	powershellArguments := " -HubSiteId " + hubId
	return RunPNP[SitePages]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
}

func GetSiteNavigation(siteURL string) (*SiteNavigation, error) {
	powershellScript := "scripts/sharepoint/get-site-navigation.ps1"
	powershellArguments := " -childSite " + siteURL
	return RunPNP[SiteNavigation]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
}

func GetHubSites() (*[]HubSite, error) {
	powershellScript := "scripts/sharepoint/get-hubsites.ps1"
	powershellArguments := ""
	return RunPNP[[]HubSite]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
}

func CopyLibrary(sourceUrl string, destinationUrl string, sourceLibray string, destinationLibray string) (*[]HubSite, error) {
	powershellScript := "scripts/sharepoint/copy-library.ps1"
	powershellArguments := fmt.Sprintf("-SourceSiteURL \"%s\" -DestinationSiteURL  \"%s\" -SourceLibraryName \"%s\" -DestinationLibraryName  \"%s\"", sourceUrl, destinationUrl, sourceLibray, destinationLibray)
	return RunPNP[[]HubSite]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
}

func CopyPage(sourceUrl string, destinationUrl string, pageName string, destpageName string) (*CopyPageResult, error) {

	powershellScript := "scripts/sharepoint/copy-page.ps1"
	powershellArguments := fmt.Sprintf("-SourceSiteURL \"%s\" -DestinationSiteURL  \"%s\" -PageName \"%s\"  -DestPageName \"copy-%s\"", sourceUrl, destinationUrl, pageName, destpageName)
	return RunPNP[CopyPageResult]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
}

func RenameLibrary(url string, fromlibrary string, tolibrary string, newUrl string) (*[]HubSite, error) {
	powershellScript := "scripts/sharepoint/rename-library.ps1"
	powershellArguments := fmt.Sprintf("-SourceSiteURL \"%s\" -oldListName  \"%s\" -newListName \"%s\" -newListUrl  \"%s\"", url, fromlibrary, tolibrary, newUrl)
	return RunPNP[[]HubSite]("koksmat", powershellScript, powershellArguments, "", CallbackMockup)
}

func GetSiteTemplate(url string) (*[]HubSite, error) {
	powershellScript := "scripts/sharepoint/get-site-template.ps1"
	powershellArguments := fmt.Sprintf(`-Url "%s"`, url)
	return RunPNP[[]HubSite]("koksmat", powershellScript, powershellArguments, "", func(workingDirectory string) {
		template, err := os.ReadFile(path.Join(workingDirectory, "template.xml"))
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s", template)

	})
}
