package restapi

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func validateSubscription(w http.ResponseWriter, r *http.Request) {
	token := r.URL.Query().Get("validationtoken")
	if token != "" {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Header().Set("X-Content-Type-Options", "nosniff")
		w.WriteHeader(200)
		fmt.Fprint(w, token)
		return
	}
	p := ListNotication{}
	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	for _, v := range p.Value {
		log.Println("Resource", v.Resource)
		log.Println("SiteURL", v.SiteURL)
		log.Println("WebID", v.WebID)
		log.Println("SubscriptionID", v.SubscriptionID)
		log.Println("ClientState", v.ClientState)
		log.Println("ExpirationDateTime", v.ExpirationDateTime)
		log.Println("TenantID", v.TenantID)

	}

}

func getBlobb(path string) (*string, error) {
	return nil, nil
}

func getBlob() usecase.Interactor {
	type BlobRequest struct {
		Tag string `json:"tag" path:"tag" example:"SITEMAP%7Chttps%3A%2F%2Fchristianiabpos.sharepoint.com%2Fsites%2Fnexiintra-home"  binding:"required"`
	}
	type BlobResponse struct {
		Content map[string]interface{} `json:"content,inline"`
		Cache   string                 `header:"Cache-Control" json:"-"`
	}

	u := usecase.NewInteractor(func(ctx context.Context, input BlobRequest, output *BlobResponse) error {

		o, err := model.GetBlob(input.Tag)
		if err != nil {
			return err
		}
		br := &BlobResponse{
			Content: *o,
			Cache:   "public, max-age=60",
		}
		*output = *br
		return err

	})

	u.SetTitle("Reading blob")
	u.SetDescription(`


Returns a piece of unstructured data 

`)
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(authenticationTag)
	return u
}
