package model

import (
	"context"
	"encoding/json"
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

func ProvisionRoomBySharePointID(id int) (*string, error) {

	sp, err := sharepoint.GetClient("https://christianiabpos.sharepoint.com/sites/Cava3")
	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", "Rooms"))

	itemID := id

	item, err := list.Items().GetByID(itemID).Get()
	if err != nil {
		log.Println("Rooms not found in SharePoint")
		return nil, err
	}

	result, err := powershell.CreateRoom(item.Data().Title, 1)
	if err != nil {
		log.Println("PowerShell returned error")
		return nil, err
	}
	newStatus := "Provisioned"
	itemUpdatePayload := []byte(fmt.Sprintf(`{
	"Provisioning_x0020_Status": "%s",
	"Email": "%s"}`, newStatus, result.MailAddress))

	_, err = list.Items().GetByID(itemID).Update(itemUpdatePayload)

	if err != nil {
		log.Println("Could not update SharePoint")
		return nil, err
	}

	return &result.MailAddress, nil

}

func DeleteRoomBySharePointID(id int) (*string, error) {

	sp, err := sharepoint.GetClient("https://christianiabpos.sharepoint.com/sites/Cava3")
	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", "Rooms"))

	itemID := id

	data, err := list.Items().Select("Id,Title,Capacity,Provisioning_x0020_Status,Email,RestrictedTo,TeamsMeetingRoom,Canbeusedforreceptions,DeviceSerialNumber,Price_x0020_List/Deliverto,CiscoVideo,Production").GetByID(itemID).Get()
	if err != nil {
		log.Println("Rooms not found in SharePoint")
		return nil, err
	}

	i := new(sharepoint.Room)
	err = json.Unmarshal(data.Normalized(), i)
	if err != nil {
		return nil, err
	}

	_, err = powershell.RemoveRoom(i.Email)
	if err != nil {
		log.Println("PowerShell returned error")
		return nil, err
	}
	newStatus := "Deleted"
	itemUpdatePayload := []byte(fmt.Sprintf(`{
	"Provisioning_x0020_Status": "%s"
	}`, newStatus))

	_, err = list.Items().GetByID(itemID).Update(itemUpdatePayload)

	if err != nil {
		log.Println("Could not update SharePoint")
		return nil, err
	}

	return nil, nil

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
	r.Room.Email = result.MailAddress
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
func RoomsToUpdate() error {
	r, err := db.FindOne[*Rooms](&Rooms{}, bson.D{{"provisioningstatus", "Update"}})
	if err != nil {
		log.Println("No rooms to update")
		return err
	}
	log.Println("Updating room", r.Room.Title)
	sp, err := sharepoint.GetClient("https://christianiabpos.sharepoint.com/sites/Cava3")
	list := sp.Web().GetList(fmt.Sprintf("Lists/%s", "Rooms"))

	itemID := r.Room.ID

	_, err = list.Items().GetByID(itemID).Get()
	if err != nil {
		log.Println("Rooms not found in SharePoint")
		return err
	}

	_, err = powershell.SetPolicy(r.Room.Email, r.Room.RestrictedTo)
	if err != nil {
		log.Println("PowerShell returned error", err)
		return err
	}
	newStatus := "Updated"
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

func SetPolicyOnAllRooms() error {
	rooms, err := db.GetFiltered[*Rooms](&Rooms{}, bson.D{{"provisioningstatus", "Provisioned"}})
	if err != nil {
		return err
	}

	for _, room := range rooms {
		log.Println("Setting policy on room", room.Room.Title)
		_, err := powershell.SetPolicy(room.Room.Email, room.Room.RestrictedTo)
		if err != nil {
			log.Println("PowerShell returned error", err)
		}
	}
	return nil
}

type RoomGroup struct {
	FloorName string `json:"floorName"`
	Rooms     []struct {
		Title    string `json:"title"`
		Floor    string `json:"floor"`
		Building string `json:"building"`
		Location string `json:"location"`
		Country  string `json:"country"`
		Mail     string `json:"mail"`
		Capacity int    `json:"capacity"`
	} `json:"rooms"`
	TotalRooms int `json:"totalRooms"`
}
type RoomGroups []RoomGroup

func RoomsGroups() (*RoomGroups, error) {
	// Requires the MongoDB Go Driver
	// https://go.mongodb.org/mongo-driver
	ctx := context.TODO()

	coll := mgm.Coll(&Rooms{})
	cursor, err := coll.Aggregate(ctx, bson.A{
		bson.D{{"$match", bson.D{{"provisioningstatus", "Provisioned"}}}},
		bson.D{
			{"$replaceRoot",
				bson.D{
					{"newRoot",
						bson.D{
							{"title", "$title"},
							{"floor", "$floor.title"},
							{"building", "$building.title"},
							{"location", "$location.title"},
							{"country", "$country.title"},
							{"mail", "$email"},
							{"capacity", "$capacity"},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$floor"},
					{"rooms", bson.D{{"$push", "$$ROOT"}}},
				},
			},
		},
		bson.D{
			{"$addFields",
				bson.D{
					{"floorName", "$_id"},
					{"totalRooms", bson.D{{"$size", "$rooms"}}},
				},
			},
		},
	})
	if err != nil {
		return nil, err
	}
	records := RoomGroups{}
	defer cursor.Close(ctx)
	for cursor.Next(ctx) {
		var record RoomGroup
		if err = cursor.Decode(&record); err != nil {
			return nil, err
		}
		if (record.FloorName != "") && (record.TotalRooms > 0) {
			records = append(records, record)
		}
	}

	return &records, nil
}

//"Id,Title,Capacity,Provisioning_x0020_Status"
