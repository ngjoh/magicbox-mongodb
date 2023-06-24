package model

import (
	"log"

	"github.com/kamva/mgm/v3"

	"github.com/koksmat-com/koksmat/sharepoint"
)

type Rooms struct {
	mgm.DefaultModel `bson:",inline"`
	sharepoint.Room  `bson:",inline"`
	sharepoint.Floor
	sharepoint.Building
	sharepoint.Locations
	sharepoint.Country
}

func ImportRooms() error {
	rooms, err := sharepoint.RoomsList()
	if err != nil {
		return err
	}
	floors, err := sharepoint.FloorsList()
	if err != nil {
		return err
	}
	buildings, err := sharepoint.BuildingsList()
	if err != nil {
		return err
	}
	locations, err := sharepoint.LocationsList()
	if err != nil {
		return err
	}
	countries, err := sharepoint.CountriesList()
	if err != nil {
		return err
	}

	// floors,err := FloorsList()
	// if err != nil {
	// 	return err
	// }

	for _, room := range *rooms {
		log.Println(room.Title)
		roomRecord := Rooms{Room: room}
		for _, floor := range *floors {
			for _, roomOnFloor := range floor.Rooms.Results {
				if roomOnFloor.ID == room.ID {
					roomRecord.Floor = floor
					break
				}
			}
		}
		for _, building := range *buildings {
			for _, floorInBuilding := range building.Floors.Results {
				if floorInBuilding.ID == roomRecord.Floor.ID {
					roomRecord.Building = building
					break
				}
			}
		}
		for _, location := range *locations {
			for _, buildingInLocation := range location.Buildings.Results {
				if buildingInLocation.ID == roomRecord.Building.ID {
					roomRecord.Locations = location
					break

				}
			}
		}

		for _, country := range *countries {
			for _, locationInCountry := range country.Locations.Results {
				if locationInCountry.ID == roomRecord.Locations.ID {
					roomRecord.Country = country
					break
				}
			}
		}

		if err := mgm.Coll(&roomRecord).Create(&roomRecord); err != nil {
			return err
		}
	}
	return nil
}

//"Id,Title,Capacity,Provisioning_x0020_Status"
