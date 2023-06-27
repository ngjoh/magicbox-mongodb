package sharepoint

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetCavaRoomList(t *testing.T) {
	items, err := RoomsList()

	if err != nil {
		t.Error(err)
	}

	if len(*items) == 0 {
		t.Error("No items")
	}
	foundRoomCapacity := false
	for _, item := range *items {
		//	t.Log(item.Title)
		if item.Capacity > 0 {
			foundRoomCapacity = true
		}
	}

	assert.True(t, foundRoomCapacity, "Expected at least one room to have capacity set")

}

func TestGetCavaFloorList(t *testing.T) {
	items, err := FloorsList()

	if err != nil {
		t.Error(err)
	}

	if len(*items) == 0 {
		t.Error("No items")
	}
	foundRooms := false
	for _, item := range *items {
		//	t.Log(item.Title)
		if len(item.Rooms.Results) > 0 {
			foundRooms = true
		}
	}

	assert.True(t, foundRooms, "Expected at least one room to be referenced")

}

func TestGetCavaBuildingList(t *testing.T) {
	items, err := BuildingsList()

	if err != nil {
		t.Error(err)
	}

	if len(*items) == 0 {
		t.Error("No items")
	}
	foundFloors := false
	for _, item := range *items {
		//	t.Log(item.Title)
		if len(item.Floors.Results) > 0 {
			foundFloors = true
		}
	}

	assert.True(t, foundFloors, "Expected at least one floor to be referenced")

}

func TestGetCavaLocationsList(t *testing.T) {
	items, err := LocationsList()

	if err != nil {
		t.Error(err)
	}

	if len(*items) == 0 {
		t.Error("No items")
	}
	foundBuildings := false
	for _, item := range *items {
		//	t.Log(item.Title)
		if len(item.Buildings.Results) > 0 {
			foundBuildings = true
		}
	}

	assert.True(t, foundBuildings, "Expected at least one building to be referenced")

}

func TestGetCavaCountriesList(t *testing.T) {
	items, err := CountriesList()

	if err != nil {
		t.Error(err)
	}

	if len(*items) == 0 {
		t.Error("No items")
	}
	foundCountryCode := false
	for _, item := range *items {
		//	t.Log(item.Title)
		if item.Countrycode != "" {
			foundCountryCode = true
		}
	}

	assert.True(t, foundCountryCode, "Expected country code to be set")

}

func TestGetChanges(t *testing.T) {
	err := GetChanges("https://christianiabpos.sharepoint.com/sites/Cava3", "Test Changes")

	assert.Nil(t, err)

}
