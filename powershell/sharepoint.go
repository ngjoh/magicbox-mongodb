package powershell

import "time"

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

func GetHubSpokesSitePages(hubId string) (*SitePages, error) {
	powershellScript := "scripts/sharepoint/get-hubsite-spokes-pages.ps1"
	powershellArguments := " -HubSiteId " + hubId
	return RunPNP[SitePages](powershellScript, powershellArguments)
}

func GetSiteNavigation(siteURL string) (*SiteNavigation, error) {
	powershellScript := "scripts/sharepoint/get-site-navigation.ps1"
	powershellArguments := " -childSite " + siteURL
	return RunPNP[SiteNavigation](powershellScript, powershellArguments)
}

func GetHubSites() (*[]HubSite, error) {
	powershellScript := "scripts/sharepoint/get-hubsites.ps1"
	powershellArguments := ""
	return RunPNP[[]HubSite](powershellScript, powershellArguments)
}
