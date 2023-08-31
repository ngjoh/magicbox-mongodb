package model

import (
	"log"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/officegraph"
)

type Device struct {
	mgm.DefaultModel `bson:",inline"`
	Device           officegraph.Device `bson:",inline"`
}

func NewDevice(device officegraph.Device) error {

	newRecord := &Device{
		DefaultModel: mgm.DefaultModel{},
		Device:       device,
	}

	return mgm.Coll(newRecord).Create(newRecord)

}

func SyncDevices() error {
	log.Println("Reading devices")
	devices, err := officegraph.GetDevices()
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Saving to collection")
	for _, device := range *devices {
		err = NewDevice(device)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	log.Println("Done")
	return nil
}
