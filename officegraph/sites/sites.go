package sites

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/koksmat-com/koksmat/shared"
	"github.com/koksmat-com/koksmat/sharepoint/sites/nexiintra_home"
)

type NewsChannelsListItem struct {
	shared.Item `bson:",inline"`
	NewsChannel *nexiintra_home.SP_NewsChannels `json:"fields,inline"`
}

func httpGet[T any](token string, url string) (result *[]T, err error) {

	type ListItem struct {
		OdataNextLink *string `json:"@odata.nextLink"`
		Value         *[]T    `json:"value,omitempty"`
	}
	nextLink := url
	records := []T{}
	for nextLink != "" {

		req, err := http.NewRequest("GET", nextLink, nil)
		req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
		client := &http.Client{}
		rsp, err := client.Do(req)

		if err != nil {
			return nil, err
		}

		if !(strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2) {
			return nil, fmt.Errorf("Invalid response")
		}

		var record ListItem
		bodyBytes, err := io.ReadAll(rsp.Body)

		fmt.Println("******** Copy to clipboard ********")
		clipboard.WriteAll(string(bodyBytes))

		defer func() { _ = rsp.Body.Close() }()
		err = json.Unmarshal(bodyBytes, &record)
		if err != nil {
			return nil, err
		}

		records = append(records, *record.Value...)

		if record.OdataNextLink == nil {
			nextLink = ""
		} else {
			nextLink = *record.OdataNextLink
		}

	}
	return &records, nil
}

func GetListItems[T any](token string, sitePath string, listName string, additionalFields string) (*[]T, error) {
	additionalFieldsWithCommaPrefix := ""
	if additionalFields != "" {
		additionalFieldsWithCommaPrefix = "," + additionalFields
	}
	endPoint := fmt.Sprintf(`https://graph.microsoft.com/v1.0/sites/christianiabpos.sharepoint.com:/%s:/lists/%s/items?$expand=fields%s`, sitePath, listName, additionalFieldsWithCommaPrefix)
	// fmt.Println("******** Copy to clipboard ********")
	// clipboard.WriteAll(endPoint)
	items, err := httpGet[T](token, endPoint)
	if err != nil {
		return nil, err
	}
	return items, nil

}
