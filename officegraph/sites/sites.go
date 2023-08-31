package sites

import (
	"fmt"

	"github.com/koksmat-com/koksmat/shared"
	"github.com/koksmat-com/koksmat/sharepoint/sites/nexiintra_home"
	"github.com/koksmat-com/koksmat/util"
)

type NewsChannelsListItem struct {
	shared.Item `bson:",inline"`
	NewsChannel *nexiintra_home.SP_NewsChannels `json:"fields,inline"`
}

func GetListItems[T any](token string, sitePath string, listName string, additionalFields string) (*[]T, error) {
	additionalFieldsWithCommaPrefix := ""
	if additionalFields != "" {
		additionalFieldsWithCommaPrefix = "," + additionalFields
	}
	endPoint := fmt.Sprintf(`https://graph.microsoft.com/v1.0/sites/christianiabpos.sharepoint.com:/%s:/lists/%s/items?$expand=fields%s`, sitePath, listName, additionalFieldsWithCommaPrefix)
	// fmt.Println("******** Copy to clipboard ********")
	// clipboard.WriteAll(endPoint)
	items, err := util.HttpGet[T](token, endPoint)
	if err != nil {
		return nil, err
	}
	return items, nil

}
