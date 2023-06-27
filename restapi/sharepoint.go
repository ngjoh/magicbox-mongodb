package restapi

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
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
