package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

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

	// log.Println(string(r.Body))
	p := &Callback{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	for _, v := range p.Value {
		model.SaveWebhookEvent(v)
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
