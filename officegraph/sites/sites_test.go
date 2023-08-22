package sites

import (
	"testing"

	"github.com/koksmat-com/koksmat/officegraph"
)

func TestGetListItems(t *testing.T) {

	_, token, err := officegraph.GetClient()
	if err != nil {
		t.Fatalf(err.Error())
	}

	got, err := GetListItems[NewsChannelsListItem](token, "sites/nexiintra-home", "News Channels", "")
	if err != nil {
		t.Fatalf(err.Error())
	}
	if (got) == nil {
		t.Fatalf("Should not return nil")
	}

}
