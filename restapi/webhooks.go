package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/koksmat-com/koksmat/model"
	"github.com/koksmat-com/koksmat/officegraph"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

const webhooksTag = "Webhooks"

type Callback struct {
	Value []model.WebhookEventStruct `json:"value"`
}

func validateSubscription(w http.ResponseWriter, r *http.Request) {

	token := r.URL.Query().Get("validationToken")
	if token != "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(200)
		fmt.Fprint(w, token)
		return
	}

	// bodyBytes, _ := io.ReadAll(r.Body)
	// defer func() { _ = r.Body.Close() }()
	// log.Println(string(bodyBytes))

	p := &Callback{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return

	}

	_, authToken, _ := officegraph.GetClient()
	for _, v := range p.Value {
		if v.ClientState == "room" {
			req, err := http.NewRequest("GET", fmt.Sprintf("https://graph.microsoft.com/v1.0/%s?$select=subject,body,bodyPreview,organizer,attendees,start,end,location", v.Resource), nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", authToken))
			client := &http.Client{}
			rsp, err := client.Do(req)

			if err != nil {
				log.Println(err)
				return
			}

			if strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2 {
				eventItem := model.EventStruct{}

				bodyBytes, err := io.ReadAll(rsp.Body)
				defer func() { _ = rsp.Body.Close() }()
				err = json.Unmarshal(bodyBytes, &eventItem)
				if err != nil {
					log.Println(err)
					return
				}

				cavaId := ""
				cavaStart := strings.Index(eventItem.Body.Content, "https://cava.nets-intranets.com")
				if cavaStart > 0 {
					cavaEnd := strings.Index(eventItem.Body.Content[cavaStart:], "\"")
					cavaId = eventItem.Body.Content[cavaStart : cavaStart+cavaEnd]
					log.Println("cava id", cavaId)
				}

				// os.WriteFile("event.json", bodyBytes, 0644)

				model.SaveWebhookUserEvent(v, eventItem, cavaId)
			} else {
				model.SaveWebhookEvent(v)
			}
		} else {
			model.SaveWebhookEvent(v)
		}

	}

	// for _, v := range p.Value {

	// 	log.Println("Resource", v.Resource)
	// 	log.Println("SiteURL", v.SiteURL)
	// 	log.Println("WebID", v.WebID)
	// 	log.Println("SubscriptionID", v.SubscriptionID)
	// 	log.Println("ClientState", v.ClientState)
	// 	log.Println("ExpirationDateTime", v.ExpirationDateTime)
	// 	log.Println("TenantID", v.TenantID)

	// }
	w.WriteHeader(200)
	fmt.Fprint(w, "received")

}
func getWebHooks() usecase.Interactor {
	type GetRequest struct {
		//	Paging `bson:",inline"`
	}

	type GetResponse struct {
		Webhooks []*officegraph.MicrosoftGraphSubscription `json:"webhooks"`
		// NumberOfRecords int64                                     `json:"numberofrecords"`
		// Pages           int64                                     `json:"pages"`
		// CurrentPage     int64                                     `json:"currentpage"`
		// PageSize        int64                                     `json:"pagesize"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := officegraph.SubscriptionList()
		output.Webhooks = data
		// output.NumberOfRecords = int64(len(data))
		// output.Pages = 1
		// output.CurrentPage = 1
		// output.PageSize = 100

		return err

	})

	u.SetTitle("Get webhooks ")

	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		webhooksTag,
	)
	return u
}
