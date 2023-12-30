package journeys

// import (
// 	"log"
// 	"testing"

// 	"github.com/stretchr/testify/assert"
// )

// func TestSail(t *testing.T) {
// 	id, err := DispatchShip("cava")
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	log.Println(id)

// 	assert.NotNil(t, id)

// }

// func TestGetJourney(t *testing.T) {
// 	cli, err := ConnectETCD()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	defer cli.Close()
// 	tripIds, err := EtcdGetKeys(*cli, "cava")

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	log.Println(tripIds)

// 	assert.NotNil(t, tripIds)

// }

// func TestSeedFirst(t *testing.T) {
// 	trip, err := GetAnyJourney()

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	log.Println(trip)

// 	need, _ := getEntity(trip.Waypoints[0].Loads.Containers[0].Needs[0])
// 	assert.NotNil(t, need)
// 	log.Println(need)

// }

// func TestAnalyseJourney(t *testing.T) {
// 	trip, err := GetAnyJourney()

// 	if err != nil {
// 		t.Error(err)
// 	}
// 	log.Println(trip)
// 	AnalyzeTrip(&trip.)
// 	assert.NotNil(t, trip)

// }

// func TestAddCargoToTrip(t *testing.T) {
// 	//journeyName := "cava"
// 	trip, err := GetAnyJourney()
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	//key := journeyKey(journeyName, trip.Id)
// 	if err != nil {
// 		t.Error(err)
// 	}
// 	need, _ := getEntity(trip.Waypoints[0].Loads.Containers[0].Needs[0])

// 	//log.Println(trip)
// 	AddCargoToTrip(&trip, need, `{"name":"test"}`)
// 	assert.NotNil(t, trip)

// }

// func TestEtchConnection(t *testing.T) {
// 	cli, err := ConnectETCD()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	defer cli.Close()
// 	assert.Nil(t, err)
// }

// func TestEtchReadWrite(t *testing.T) {
// 	cli, err := ConnectETCD()
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	defer cli.Close()
// 	err = EtcdPut(*cli, "x", "y")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}
// 	value, err := EtcdGetFirst(*cli, "x")
// 	if err != nil {
// 		t.Error(err)
// 		return
// 	}

// 	assert.Equal(t, "y", value)
// }
