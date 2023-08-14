package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/sharepoint"
)

type Subscription struct {
	mgm.DefaultModel `bson:",inline"`
	Name             string `json:"Name"`
	SubscriptionID   string `json:"subscriptionID"`
	Type             string `json:"Type"` // "SharePoint:List", "Exchange:Mailbox", "Exchange:SharedMailbox", "Exchange:Room
	URL              string `json:"url"`
	Listname         string `json:"listname"`
}

func GetSubscriptions(name string) ([]*Subscription, error) {
	return db.GetAll[*Subscription](&Subscription{})
}

func NewSubscription(name string, sharePointSiteUrl string, listname string, callbackUrl string) error {
	s, err := sharepoint.CreateSubscription(sharePointSiteUrl, listname, callbackUrl, time.Now().Add(time.Hour*10), "xx")
	if err != nil {
		return err
	}

	newRecord := &Subscription{
		DefaultModel:   mgm.DefaultModel{},
		Name:           name,
		SubscriptionID: s.ID,
		URL:            sharePointSiteUrl,
		Listname:       listname,
		Type:           "SharePoint:List",
	}
	return mgm.Coll(newRecord).Create(newRecord)
}
