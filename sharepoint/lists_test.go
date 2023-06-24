package sharepoint

import (
	"testing"

	"github.com/koksmat-com/koksmat/exchange/rooms"
)

func TestGetCavaRoomList(t *testing.T) {
	items, err := GetListItems[rooms.Room]("https://christianiabpos.sharepoint.com/sites/Cava3", "Rooms", "Id,Title,Capacity,Provisioning_x0020_Status")

	if err != nil {
		t.Error(err)
	}

	if len(items) == 0 {
		t.Error("No items")
	}

}
