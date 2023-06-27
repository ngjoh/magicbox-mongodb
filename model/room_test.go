package model

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRoomsToProvision(t *testing.T) {
	RoomsToProvision()

	// /	assert.Nil(t, err)
}

func TestRoomsToDelete(t *testing.T) {
	RoomsToDelete()

	// /	assert.Nil(t, err)
}

func TestPolicy(t *testing.T) {
	SetPolicyOnAllRooms()
}

func TestGetAllRooms(t *testing.T) {
	x, err := RoomsGroups()

	assert.Nil(t, err)

	for _, floor := range *x {
		log.Println("----", floor.FloorName, "-----------")
		for _, room := range floor.Rooms {
			log.Println(room.Title)
		}
	}

}
