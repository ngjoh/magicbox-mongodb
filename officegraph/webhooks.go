package officegraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
)

func RemoveSubscription(id string) (*http.Response, error) {
	c, _, err := GetClient() //NewClient("https://graph.microsoft.com/v1.0/")
	if err != nil {
		return nil, err
	}
	ctx := context.Background()

	return c.SubscriptionsSubscriptionDeleteSubscription(ctx, id, &SubscriptionsSubscriptionDeleteSubscriptionParams{})
}
func SubscriptionList() (items []*MicrosoftGraphSubscription, err error) {
	c, token, err := GetClient() //NewClient("https://graph.microsoft.com/v1.0/")
	if err != nil {
		return
	}
	ctx := context.Background()

	SubscriptionsSubscriptionListSubscriptionParams := &SubscriptionsSubscriptionListSubscriptionParams{}
	response, err := c.SubscriptionsSubscriptionListSubscriptionWithResponse(ctx, SubscriptionsSubscriptionListSubscriptionParams)
	if err != nil {
		return nil, err
	}

	hasMore := true
	nextLink := response.JSON2XX.OdataNextLink
	values := *response.JSON2XX.Value
	result := []*MicrosoftGraphSubscription{}
	for hasMore {

		for _, v := range values {
			result = append(result, &v)
		}

		if nextLink == nil {
			hasMore = false
		} else {
			//log.Println(*nextLink)
			req, err := http.NewRequest("GET", *nextLink, nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", token))
			client := &http.Client{}
			rsp, err := client.Do(req)

			if err != nil {
				return nil, err
			}

			if strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2 {
				var dest MicrosoftGraphSubscriptionCollectionResponse
				bodyBytes, err := io.ReadAll(rsp.Body)
				defer func() { _ = rsp.Body.Close() }()
				err = json.Unmarshal(bodyBytes, &dest)
				if err != nil {
					return nil, err
				}
				values = *dest.Value
				nextLink = dest.OdataNextLink
			} else {
				hasMore = false
			}

		}

	}
	return result, nil
}
