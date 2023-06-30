package model

import (
	"log"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/powershell"
	"go.mongodb.org/mongo-driver/bson"
)

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
type Site struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string     `json:"name"`
	SiteURL          string     `json:"siteURL"`
	HomePage         string     `json:"homePage"`
	Navigation       Navigation `json:"navigation"`
	Pages            Pages      `json:"pages"`
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
		changedRecord, err := db.FindOne[*Site](&Site{}, filter)

		if err != nil {

			newRecord := &Site{
				DefaultModel: mgm.DefaultModel{},
				Name:         site.Title,
				SiteURL:      site.Siteurl,
				HomePage:     site.WelcomePage,

				Pages: pages,
			}

			mgm.Coll(newRecord).Create(newRecord)

		} else {

			changedRecord.HomePage = site.WelcomePage
			changedRecord.Pages = pages
			changedRecord.Name = site.Title

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

	sites, err := db.GetAll[*Site](&Site{})
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
		changedRecord, err := db.FindOne[*Site](&Site{}, filter)

		if err != nil {

			newRecord := &Site{
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
