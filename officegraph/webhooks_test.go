package officegraph

import (
	"context"
	"fmt"
	"io"
	"log"
	"testing"
	"time"
)

func TestAuth(t *testing.T) {
	_, _, err := GetClient()
	if err != nil {
		t.Fatalf("Should not return error")
	}
}
func TestList3(t *testing.T) {

	items, err := SubscriptionList()
	if err != nil {
		t.Fatalf("Should not return error")
	}
	log.Println(len(items), "Items")
	for _, item := range items {
		log.Println(*item.Resource)
	}
}

func TestRemover2(t *testing.T) {
	items, err := SubscriptionList()
	if err != nil {
		t.Fatalf("Should not return error")
	}

	fmt.Println("Removing", len(items), "items")
	x := len(items)
	for _, item := range items {

		log.Println("Removing", *item.Resource, x, "remaining")
		_, err = RemoveSubscription(*item.Id)
		if err != nil {
			log.Println("Removing", *item.Resource)
			t.Fatalf("Should not return error")
		}
		x--
	}

}

func s(text string) *string {
	return &text
}
func TestAddRoomDevNiels(t *testing.T) {
	c, _, err := GetClient() //NewClient("https://graph.microsoft.com/v1.0/")
	if err != nil {
		t.Fatalf("Should not return error")
	}
	ctx := context.Background()

	time := time.Now().Add(time.Hour * 24 * 1)

	sub := &MicrosoftGraphSubscription{
		//Id:                 s("1"),
		ChangeType: s("created,updated,deleted"),
		//Resource:           s("https://christianiabpos.sharepoint.com/sites/Cava3/_api/Web/GetList('/sites/Cava3/Lists/Test Changes"),
		Resource:           s("/users/room-dk-kb601-31m3@nets.eu/events"),
		ExpirationDateTime: &time,

		//NotificationUrl: s("https://niels-mac.nets-intranets.com/api/v1/subscription/notify"),
		NotificationUrl: s("https://magicbox.nexi-intra.com/api/v1/subscription/notify"),
		ClientState:     s("room"),
	}

	response, err := c.SubscriptionsSubscriptionCreateSubscription(ctx, *sub)

	if err != nil {
		t.Fatalf("Should not return error")
	}
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
func TestAddUsersDevNiels(t *testing.T) {
	c, _, err := GetClient() //NewClient("https://graph.microsoft.com/v1.0/")
	if err != nil {
		t.Fatalf("Should not return error")
	}
	ctx := context.Background()

	time := time.Now().Add(time.Hour * 24 * 1)

	sub := &MicrosoftGraphSubscription{
		//Id:                 s("1"),
		ChangeType: s("created,updated,deleted"),
		//Resource:           s("https://christianiabpos.sharepoint.com/sites/Cava3/_api/Web/GetList('/sites/Cava3/Lists/Test Changes"),
		Resource:           s("/users"),
		ExpirationDateTime: &time,

		NotificationUrl: s("https://niels-mac.nets-intranets.com/api/v1/subscription/notify"),
		//NotificationUrl: s("https://magicbox.nexi-intra.com/api/v1/subscription/notify"),
		ClientState: s("zz"),
	}

	response, err := c.SubscriptionsSubscriptionCreateSubscription(ctx, *sub)

	if err != nil {
		t.Fatalf("Should not return error")
	}
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
}

func TestAddRoomProd(t *testing.T) {
	c, _, err := GetClient() //NewClient("https://graph.microsoft.com/v1.0/")
	if err != nil {
		t.Fatalf("Should not return error")
	}
	ctx := context.Background()

	time := time.Now().Add(time.Hour * 24 * 2)

	sub := &MicrosoftGraphSubscription{
		//Id:                 s("1"),
		ChangeType: s("created,updated,deleted"),
		//Resource:           s("https://christianiabpos.sharepoint.com/sites/Cava3/_api/Web/GetList('/sites/Cava3/Lists/Test Changes"),
		Resource:           s("/users/room-dk-kb601-31m3@nets.eu/events"),
		ExpirationDateTime: &time,

		NotificationUrl: s("https://magicbox.nexi-intra.com/api/v1/subscription/notify"),
		ClientState:     s("room"),
	}

	response, err := c.SubscriptionsSubscriptionCreateSubscription(ctx, *sub)

	if err != nil {
		t.Fatalf("Should not return error")
	}
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
func TestAddUsersProd(t *testing.T) {
	c, _, err := GetClient() //NewClient("https://graph.microsoft.com/v1.0/")
	if err != nil {
		t.Fatalf("Should not return error")
	}
	ctx := context.Background()

	time := time.Now().Add(time.Hour * 24 * 2)

	sub := &MicrosoftGraphSubscription{
		//Id:                 s("1"),
		ChangeType: s("created,updated,deleted"),
		//Resource:           s("https://christianiabpos.sharepoint.com/sites/Cava3/_api/Web/GetList('/sites/Cava3/Lists/Test Changes"),
		Resource:           s("/users"),
		ExpirationDateTime: &time,

		NotificationUrl: s("https://magicbox.nexi-intra.com/api/v1/subscription/notify"),
		ClientState:     s("zz"),
	}

	response, err := c.SubscriptionsSubscriptionCreateSubscription(ctx, *sub)

	if err != nil {
		t.Fatalf("Should not return error")
	}
	body, err := io.ReadAll(response.Body)
	fmt.Println(string(body))
}
