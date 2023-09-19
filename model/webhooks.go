package model

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/config"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/officegraph"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	Processed          bool   `json:"processed"`
	// EventStruct        `bson:"event"`
	// CavaId             string `json:"cavaId"`
}

func GetWebhookss(name string) ([]*Subscription, error) {
	return db.GetAll[*Subscription](&Subscription{})
}

func SaveWebhookEvent(data WebhookEventStruct) error {

	newRecord := &WebhookEvent{
		DefaultModel:       mgm.DefaultModel{},
		Version:            "2",
		WebhookEventStruct: data,
	}
	return mgm.Coll(newRecord).Create(newRecord)
}

func (e *WebhookEvent) Hash() string {
	return fmt.Sprintf("%s:%s", e.ResourceData.OdataID, e.ResourceData.OdataEtag)
}

type CachedWebhookItems struct {
	Event WebhookEvent `bson:",inline"`
	Data  interface{}  `bson:"data"`
}

func addToCollection(databaseName string, collectionName string, event WebhookEvent, data []byte) {
	ctx := context.TODO()

	// Set client options
	clientOptions := options.Client().ApplyURI(config.MongoConnectionString())

	// Connect to MongoDB
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := client.Disconnect(ctx); err != nil {
			log.Println("addToCollection", err)
		}
	}()

	coll := client.Database(databaseName).Collection(collectionName)
	var v interface{}

	json.Unmarshal(data, &v)
	cachedItem := &CachedWebhookItems{
		Event: event,
		Data:  v,
	}

	_, err = coll.InsertOne(ctx, cachedItem)

	if err != nil {
		log.Fatal(err)
	}
}
func RunWebhookEventParser() error {
	loops := 0
	log.Println("Webhook event processor starting")
	for {
		loops++

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

		defer cancel()

		results, err := mgm.Coll(&WebhookEvent{}).Find(context.TODO(), bson.D{{"processed", false}}, options.Find().SetSort(bson.D{{"created_at", 1}}))
		if err != nil {

			return err
		}

		accessToken := ""

		for results.Next(ctx) {
			var webhookEvent WebhookEvent

			if err = results.Decode(&webhookEvent); err != nil {
				return err
			}

			if accessToken == "" {
				_, accessToken, err = officegraph.GetClient()
				if err != nil {
					log.Println("Error getting access token", err)
					continue
				}
			}
			url := fmt.Sprintf("https://graph.microsoft.com/v1.0/%s", webhookEvent.WebhookEventStruct.Resource)
			req, err := http.NewRequest("GET", url, nil)
			req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", accessToken))
			client := &http.Client{}
			rsp, err := client.Do(req)

			if err != nil {
				log.Println("Error getting resource data", err)
				continue
			}

			if strings.Contains(rsp.Header.Get("Content-Type"), "json") && rsp.StatusCode/100 == 2 {

				bodyBytes, _ := io.ReadAll(rsp.Body)
				defer func() { _ = rsp.Body.Close() }()
				//data := fmt.Sprintf("%s", bodyBytes)

				addToCollection(config.DatabaseName(), "cached_webhook_items", webhookEvent, bodyBytes)

				//log.Println(data)

			} else {
				log.Println("Error getting resource data", rsp.StatusCode)
			}
			webhookEvent.Processed = true
			updateErr := mgm.Coll(&webhookEvent).Update(&webhookEvent)
			if updateErr != nil {
				log.Println("Error updating web hook", updateErr)
			}
			//log.Println(webhookEvent.ID)

		}
		results.Close(ctx)

		time.Sleep(2 * time.Second)
		if (loops % 10) == 0 {
			log.Println("Still here", loops)
		}

	}
}
