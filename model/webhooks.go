package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
)

type EventStruct struct {
	OdataContext string `json:"@odata.context"`
	OdataEtag    string `json:"@odata.etag"`
	ID           string `json:"id"`
	Subject      string `json:"subject"`
	BodyPreview  string `json:"bodyPreview"`
	Body         struct {
		ContentType string `json:"contentType"`
		Content     string `json:"content"`
	} `json:"body"`
	Start struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"start"`
	End struct {
		DateTime string `json:"dateTime"`
		TimeZone string `json:"timeZone"`
	} `json:"end"`
	Location struct {
		DisplayName  string `json:"displayName"`
		LocationURI  string `json:"locationUri"`
		LocationType string `json:"locationType"`
		UniqueID     string `json:"uniqueId"`
		UniqueIDType string `json:"uniqueIdType"`
		Address      struct {
			Street          string `json:"street"`
			City            string `json:"city"`
			State           string `json:"state"`
			CountryOrRegion string `json:"countryOrRegion"`
			PostalCode      string `json:"postalCode"`
		} `json:"address"`
		Coordinates struct {
		} `json:"coordinates"`
	} `json:"location"`
	Attendees []struct {
		Type   string `json:"type"`
		Status struct {
			Response string    `json:"response"`
			Time     time.Time `json:"time"`
		} `json:"status"`
		EmailAddress struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"emailAddress"`
	} `json:"attendees"`
	Organizer struct {
		EmailAddress struct {
			Name    string `json:"name"`
			Address string `json:"address"`
		} `json:"emailAddress"`
	} `json:"organizer"`
	CalendarOdataAssociationLink string `json:"calendar@odata.associationLink"`
	CalendarOdataNavigationLink  string `json:"calendar@odata.navigationLink"`
}
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
	Version            string `json:"version"`
	EventStruct        `bson:"event"`
	CavaId             string `json:"cavaId"`
}

func GetWebhookss(name string) ([]*Subscription, error) {
	return db.GetAll[*Subscription](&Subscription{})
}

func SaveWebhookUserEvent(data WebhookEventStruct, userevents EventStruct, cavaId string) error {

	newRecord := &WebhookEvent{
		DefaultModel:       mgm.DefaultModel{},
		Version:            "1",
		WebhookEventStruct: data,
		EventStruct:        userevents,
		CavaId:             cavaId,
	}
	return mgm.Coll(newRecord).Create(newRecord)
}

func SaveWebhookEvent(data WebhookEventStruct) error {

	newRecord := &WebhookEvent{
		DefaultModel:       mgm.DefaultModel{},
		Version:            "1",
		WebhookEventStruct: data,
	}
	return mgm.Coll(newRecord).Create(newRecord)
}
