package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
)

type WebhookEventStruct struct {
	SubscriptionID                 string    `json:"subscriptionId"`
	SubscriptionExpirationDateTime time.Time `json:"subscriptionExpirationDateTime"`
	ChangeType                     string    `json:"changeType"`
	Resource                       string    `json:"resource"`
	ResourceData                   struct {
		OdataType string `json:"@odata.type"`
		OdataID   string `json:"@odata.id"`
		OdataEtag string `json:"@odata.etag"`
		ID        string `json:"id"`
	} `json:"resourceData"`
	ClientState string `json:"clientState"`
	TenantID    string `json:"tenantId"`
}
type WebhookEvent struct {
	mgm.DefaultModel   `bson:",inline"`
	WebhookEventStruct `bson:",inline"`
}

func GetWebhookss(name string) ([]*Subscription, error) {
	return db.GetAll[*Subscription](&Subscription{})
}

func SaveWebhookEvent(data WebhookEventStruct) error {

	newRecord := &WebhookEvent{
		DefaultModel:       mgm.DefaultModel{},
		WebhookEventStruct: data,
	}
	return mgm.Coll(newRecord).Create(newRecord)
}
