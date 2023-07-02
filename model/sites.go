package model

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/powershell"
	"go.mongodb.org/mongo-driver/bson"
)

type SharepointHubSite struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"name"`
	SiteURL          string `json:"siteURL"`
	ID               string `json:"id"`
}

type SharepointHubSites []SharepointHubSite
type Page struct {
	Editor      string `json:"Editor"`
	RelativeURL string `json:"RelativeURL"`

	CreatedOn  time.Time `json:"CreatedOn"`
	PageName   string    `json:"PageName"`
	ModifiedOn time.Time `json:"ModifiedOn"`
}
type Pages []Page

type NavigationNode struct {
	Childs      []NavigationNode `json:"childs,omitempty"`
	Title       string           `json:"title"`
	RelativeURL string           `json:"relativeURL"`
}

type Navigation []NavigationNode
type SharepointSite struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string     `json:"name"`
	SiteURL          string     `json:"siteURL"`
	HomePage         string     `json:"homePage"`
	HubSiteId        string     `json:"hubSiteId"`
	Navigation       Navigation `json:"navigation"`
	Pages            Pages      `json:"pages"`
}

func SyncHubSites() error {

	hubSites, err := powershell.GetHubSites()
	if err != nil {
		log.Println(err)
		return err
	}

	update := func(record *SharepointHubSite, src powershell.HubSite) error {
		record.Name = src.Title
		record.SiteURL = src.SiteURL

		return nil
	}
	create := func(src powershell.HubSite) (*SharepointHubSite, error) {
		record := SharepointHubSite{
			DefaultModel: mgm.DefaultModel{},
			Name:         src.Title,
			SiteURL:      src.SiteURL,
			ID:           src.ID,
		}

		return &record, nil
	}
	filter := func(src powershell.HubSite) bson.D {
		return bson.D{{"id", src.ID}}
	}
	err = db.Sync[*SharepointHubSite, powershell.HubSite](hubSites,
		update, create, filter)

	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func SyncHubSitePages(HubSiteID string) error {

	sitePages, err := powershell.GetHubSpokesSitePages(HubSiteID)
	if err != nil {
		return err
	}

	for _, site := range *sitePages {

		pages := []Page{}
		for _, page := range site.Pages {
			p := Page{
				Editor:      page.Editor,
				RelativeURL: page.RelativeURL,
				CreatedOn:   page.CreatedOn,
				PageName:    page.PageName,
				ModifiedOn:  page.ModifiedOn,
			}
			pages = append(pages, p)

		}

		filter := bson.D{{"siteurl", site.Siteurl}}
		changedRecord, err := db.FindOne[*SharepointSite](&SharepointSite{}, filter)

		if err != nil {

			newRecord := &SharepointSite{
				DefaultModel: mgm.DefaultModel{},
				Name:         site.Title,
				SiteURL:      site.Siteurl,
				HomePage:     site.WelcomePage,
				HubSiteId:    HubSiteID,
				Pages:        pages,
			}

			mgm.Coll(newRecord).Create(newRecord)

		} else {

			changedRecord.HomePage = site.WelcomePage
			changedRecord.Pages = pages
			changedRecord.Name = site.Title
			changedRecord.HubSiteId = HubSiteID
			mgm.Coll(changedRecord).Update(changedRecord)

		}

	}
	return nil
}

func iterateNodes(node powershell.NavigationNode) []NavigationNode {

	nodes := []NavigationNode{}
	for _, sharePointNode := range node.Childs {

		n1 := NavigationNode{
			Title:       sharePointNode.Title,
			RelativeURL: sharePointNode.RelativeURL,
		}
		n1.Childs = append(n1.Childs, iterateNodes(sharePointNode)...)
		nodes = append(nodes, n1)
	}

	return nodes
}
func SyncSitesNavigation() error {

	sites, err := db.GetAll[*SharepointSite](&SharepointSite{})
	if err != nil {
		return err
	}

	for _, s := range sites {
		log.Println("Syncing Site Navigation for ", s.SiteURL)
		site, err := powershell.GetSiteNavigation(s.SiteURL)

		nodes := []NavigationNode{}
		if err != nil {

			log.Println("Error  ", err)

		} else {
			for _, node := range site.Navigation {
				p := NavigationNode{
					Childs:      iterateNodes(node),
					Title:       node.Title,
					RelativeURL: node.RelativeURL,
				}
				nodes = append(nodes, p)

			}
		}
		filter := bson.D{{"siteurl", site.Siteurl}}
		changedRecord, err := db.FindOne[*SharepointSite](&SharepointSite{}, filter)

		if err != nil {

			newRecord := &SharepointSite{
				DefaultModel: mgm.DefaultModel{},
				Name:         site.Title,
				SiteURL:      site.Siteurl,
				Navigation:   nodes,
			}

			mgm.Coll(newRecord).Create(newRecord)

		} else {

			changedRecord.Navigation = nodes

			mgm.Coll(changedRecord).Update(changedRecord)

		}

	}
	return nil
}

func hasChildWithLinks(node NavigationNode) bool {
	for _, child := range node.Childs {
		if child.RelativeURL != "http://linkless.header/" {
			return true
		} else {
			return hasChildWithLinks(child)
		}
	}
	return false
}
func isSiteLink(node NavigationNode) (bool, Navigation) {
	site, err := db.FindOne[*SharepointSite](
		&SharepointSite{},
		bson.D{{"siteurl", node.RelativeURL}},
	)
	if err != nil {
		return false, nil
	} else {
		return true, site.Navigation
	}

}

func isAlreadyInHistory(siteHistory []string, siteURL string) bool {
	for _, site := range siteHistory {
		if site == siteURL {
			return true
		}
	}
	return false
}

type SiteMap struct {
	Title string     `json:"title"`
	URL   string     `json:"url"`
	Menu  []MenuItem `json:"menu"`
}
type MenuItem struct {
	Childs      []MenuItem `json:"childs,omitempty"`
	Title       string     `json:"title"`
	IsLink      bool       `json:"isLink"`
	RelativeURL string     `json:"relativeURL"`
}

/*
*

  - Recursively build the navigation tree, respecting the following rules:

  - 1. If the node has a link, add it to the tree

  - 2. If the node has no link, but has children with links, add it to the tree

  - 3. If the node has no link, and no children with links, do not add it to the tree

    If a node is linking to a site, then the navigation of the site is added to the tree, respecting the above rules. But only once per navigation structure, to avoid infinite recursion.
*/
func childNav(level int, childs []NavigationNode, siteHistory []string) []MenuItem {
	items := []MenuItem{}
	for _, child := range childs {
		if child.RelativeURL == "" {
			continue
		}
		isLink := child.RelativeURL != "http://linkless.header/"
		if isLink {
			isLinkToSite, linkedSiteNavigation := isSiteLink(child)
			if isLinkToSite {
				if isAlreadyInHistory(siteHistory, child.RelativeURL) {
					items = append(items, MenuItem{
						Childs:      []MenuItem{},
						Title:       child.Title,
						IsLink:      isLink,
						RelativeURL: child.RelativeURL,
					})
					fmt.Println(level, strings.Repeat(" ", level*4), "-", child.Title, " (Recursive site link)")

				} else {
					fmt.Println(level, strings.Repeat(" ", level*4), child.Title)
					siteHistory = append(siteHistory, child.RelativeURL)
					items = append(items, MenuItem{
						Childs:      childNav(level+1, linkedSiteNavigation, siteHistory),
						Title:       child.Title,
						IsLink:      isLink,
						RelativeURL: child.RelativeURL,
					})

				}
			} else {
				fmt.Println(level, strings.Repeat(" ", level*4), "-", child.Title)
				items = append(items, MenuItem{
					Childs:      []MenuItem{},
					Title:       child.Title,
					IsLink:      isLink,
					RelativeURL: child.RelativeURL,
				})
			}

		} else {
			hasChilds := hasChildWithLinks(child)
			if hasChilds {
				fmt.Println(level, strings.Repeat(" ", level*4), "-", child.Title, " (No link - but has childs)")
				items = append(items, MenuItem{
					Childs:      childNav(level+1, child.Childs, siteHistory),
					Title:       child.Title,
					IsLink:      isLink,
					RelativeURL: child.RelativeURL,
				})
			} else {
				fmt.Println(level, strings.Repeat(" ", level*4), "-", child.Title, " (Linkless - skipped)")

			}
		}

	}
	return items
}
func SharePointHubSitemap(homeSiteUrl string) error {

	site, err := db.FindOne[*SharepointSite](
		&SharepointSite{},
		bson.D{{"siteurl", homeSiteUrl}},
	)

	if err != nil {
		errors.New("Site not found")

	}

	sites, err := db.GetFiltered[*SharepointSite](
		&SharepointSite{},
		bson.D{{"hubsiteid", site.HubSiteId}},
	)
	if err != nil {
		return err
	}

	for _, site := range sites {
		fmt.Println("****", site.Name, "****")
		childNav(1, site.Navigation, []string{site.SiteURL})

	}
	return nil
}
func SharePointSitemap(homeSiteUrl string) (*SiteMap, error) {

	site, err := db.FindOne[*SharepointSite](
		&SharepointSite{},
		bson.D{{"siteurl", homeSiteUrl}},
	)

	if err != nil {
		return nil, errors.New("Site not found")

	}
	siteMap := SiteMap{
		Title: site.Name,
		URL:   site.SiteURL,
		Menu:  childNav(1, site.Navigation, []string{site.SiteURL}),
	}

	return &siteMap, nil
}
