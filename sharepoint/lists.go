package sharepoint

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/atotto/clipboard"
	"github.com/koltyakov/gosip/api"
)

func GetListItems[T any](sharePointSiteUrl string, listName string, selectedFields string, expand string) ([]T, error) {
	sp, err := GetClient(sharePointSiteUrl)
	if err != nil {
		return nil, err
	}

	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", listName))

	itemsResp, err := list.Items().
		Select(selectedFields). // OData $select modifier, limit what props are retrieved
		OrderBy("Id", true).    // OData $orderby modifier, defines sort order
		Expand(expand).         // OData $expand modifier, expands lookup fields
		Top(2000).              // OData $top modifier, limits page size
		Get()                   // Finalizes API constructor and sends a response

	if err != nil {
		return nil, err
	}
	results := []T{}
	// Data() method is a helper which unmarshals generic structure
	// use custom structs and unmarshal for custom fields
	for _, item := range itemsResp.Data() {
		i := new(T)
		clipboard.WriteAll(fmt.Sprintf("%s", item))
		err = json.Unmarshal(item, i)
		if err != nil {
			return nil, err
		}
		//	itemData := item.Data()
		results = append(results, *i)

		//fmt.Println(fmt.Sprintf("%s", item))

	}

	return results, nil

}

func GetSubscriptions(sharePointSiteUrl string, listName string) error {
	sp, err := GetClient(sharePointSiteUrl)
	if err != nil {
		return err
	}

	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", listName))

	subscriptions, err := list.Subscriptions().Get()
	if err != nil {
		return err
	}
	for _, item := range subscriptions {
		fmt.Println(item)
	}

	return nil

}

func CreateSubscription(sharePointSiteUrl string, listName string, notificationUrl string, expiration time.Time, clientState string) (*api.SubscriptionInfo, error) {
	sp, err := GetClient(sharePointSiteUrl)
	if err != nil {
		return nil, err
	}
	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", listName))

	subscriptions, err := list.Subscriptions().Add(notificationUrl, expiration, clientState)
	if err != nil {
		return nil, err
	}
	log.Println("Subscription add", subscriptions)
	return subscriptions, nil
}

func GetChanges(sharePointSiteUrl string, listName string) error {
	sp, err := GetClient(sharePointSiteUrl)
	if err != nil {
		return err
	}
	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", listName))
	listChangeToken, _ := list.Changes().GetCurrentToken()
	changes, err := list.Changes().GetChanges(&api.ChangeQuery{
		ChangeTokenStart: listChangeToken,

		List: true,
		Item: true,
		Add:  true,
	})
	if err != nil {
		return err
	}

	for _, change := range changes.Data() {
		fmt.Printf("%+v\n", change)
	}
	return nil
}
