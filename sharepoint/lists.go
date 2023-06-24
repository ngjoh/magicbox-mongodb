package sharepoint

import (
	"encoding/json"
	"fmt"
)

func GetListItems[T any](sharePointSiteUrl string, listName string, selectedFields string) ([]T, error) {
	sp, err := GetClient(sharePointSiteUrl)
	if err != nil {
		return nil, err
	}

	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", listName))

	itemsResp, err := list.Items().
		Select(selectedFields). // OData $select modifier, limit what props are retrieved
		OrderBy("Id", true).    // OData $orderby modifier, defines sort order
		//Expand("Provisioning_x0020_Status").                   // OData $expand modifier, expands lookup fields
		Top(1). // OData $top modifier, limits page size
		Get()   // Finalizes API constructor and sends a response

	if err != nil {
		return nil, err
	}
	results := []T{}
	// Data() method is a helper which unmarshals generic structure
	// use custom structs and unmarshal for custom fields
	for _, item := range itemsResp.Data() {
		i := new(T)
		err = json.Unmarshal(item, i)
		if err != nil {
			return nil, err
		}
		//	itemData := item.Data()
		results = append(results, *i)

	}

	return results, nil

}
