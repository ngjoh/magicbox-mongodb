package model

import (
	"fmt"
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/powershell"
	"github.com/koksmat-com/koksmat/sharepoint"
	"go.mongodb.org/mongo-driver/bson"
)

type Rooms struct {
	mgm.DefaultModel `bson:",inline"`
	sharepoint.Room  `bson:",inline"`
	sharepoint.Floor
	sharepoint.Building
	sharepoint.Location
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
					roomRecord.Location = location
					break

				}
			}
		}

		for _, country := range *countries {
			for _, locationInCountry := range country.Locations.Results {
				if locationInCountry.ID == roomRecord.Location.ID {
					roomRecord.Country = country
					break
				}
			}
		}

		_, err := db.FindOne[*Rooms](&Rooms{}, bson.D{{"id", room.ID}})

		if err != nil {
			newRecord := &roomRecord
			mgm.Coll(newRecord).Create(newRecord)
			log.Println("new")
		}

	}
	return nil
}

func RoomsToProvision() error {
	r, err := db.FindOne[*Rooms](&Rooms{}, bson.D{{"provisioningstatus", "Provision"}})
	if err != nil {
		log.Println("No rooms to provision")
		return err
	}
	log.Println("Provisioning room", r.Room.Title)
	sp, err := sharepoint.GetClient("https://christianiabpos.sharepoint.com/sites/Cava3")
	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", "Rooms"))

	itemID := r.Room.ID

	_, err = list.Items().GetByID(itemID).Get()
	if err != nil {
		log.Println("Rooms not found in SharePoint")
		return err
	}

	result, err := powershell.CreateRoom(r.Room.Title, r.Room.Capacity)
	if err != nil {
		log.Println("PowerShell returned error")
		return err
	}
	newStatus := "Provisioned"
	itemUpdatePayload := []byte(fmt.Sprintf(`{
	"Provisioning_x0020_Status": "%s",
	"Email": "%s"}`, newStatus, result.MailAddress))

	_, err = list.Items().GetByID(itemID).Update(itemUpdatePayload)

	if err != nil {
		log.Println("Could not update SharePoint")
		return err
	}

	r.Room.ProvisioningStatus = newStatus
	return mgm.Coll(r).Update(r)

}
func RoomsToDelete() error {
	r, err := db.FindOne[*Rooms](&Rooms{}, bson.D{{"provisioningstatus", "Delete"}})
	if err != nil {
		log.Println("No rooms to delete")
		return err
	}
	log.Println("Deleting room", r.Room.Title)
	sp, err := sharepoint.GetClient("https://christianiabpos.sharepoint.com/sites/Cava3")
	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", "Rooms"))

	itemID := r.Room.ID

	_, err = list.Items().GetByID(itemID).Get()
	if err != nil {
		log.Println("Rooms not found in SharePoint")
		return err
	}

	_, err = powershell.RemoveRoom(r.Room.Email)
	if err != nil {
		log.Println("PowerShell returned error", err)
		return err
	}
	newStatus := "Deleted"
	itemUpdatePayload := []byte(fmt.Sprintf(`{
	"Provisioning_x0020_Status": "%s"
	}`, newStatus))

	_, err = list.Items().GetByID(itemID).Update(itemUpdatePayload)

	if err != nil {
		log.Println("Could not update SharePoint")
		return err
	}

	r.Room.ProvisioningStatus = newStatus
	return mgm.Coll(r).Update(r)

}

//"Id,Title,Capacity,Provisioning_x0020_Status"
