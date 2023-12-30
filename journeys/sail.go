package journeys

// import (
// 	"encoding/json"
// 	"fmt"
// 	"os"

// 	"github.com/kamva/mgm/v3"

// 	"go.mongodb.org/mongo-driver/bson"
// )

// func LoadTravelPlan(journeyName string) (*TravelPlan, error) {
// 	data, err := os.ReadFile(journeyName + ".json")

// 	if err != nil {
// 		return nil, err
// 	}

// 	travelPlan := TravelPlan{}
// 	err = json.Unmarshal(data, &travelPlan)
// 	if err != nil {
// 		return nil, err
// 	}
// 	return &travelPlan, nil

// }

// func CalculateDetailedJourney(travelPlan *TravelPlan) (*DetailedJourney, error) {
// 	d := DetailedJourney{}
// 	d.Metadata.Name = travelPlan.Metadata.Name
// 	d.Metadata.Description = travelPlan.Metadata.Description
// 	d.Journey = travelPlan.Journey
// 	d.Triggers = travelPlan.Triggers

// 	for _, waypoint := range travelPlan.Waypoints {
// 		wp := WaypointDetails{}
// 		wp.Port = waypoint.Port
// 		wp.Done = waypoint.Done
// 		l := Loads{}
// 		for _, container := range waypoint.Loads.Containers {
// 			cd := ContainerDetails{}
// 			cd.Container = container.Container
// 			cd.Name = container.Name
// 			cd.Who = container.Who
// 			cd.Approve = container.Approve
// 			cd.Consult = container.Consult
// 			cd.Inform = container.Inform
// 			cd.Script = container.Script
// 			for _, need := range container.Needs {
// 				entity, entityType := getEntity(need)
// 				cd.Needs = append(cd.Needs, Entity{Name: entity, Type: entityType})
// 			}
// 			for _, produce := range container.Produces {
// 				entity, entityType := getEntity(produce)
// 				cd.Produces = append(cd.Produces, Entity{Name: entity, Type: entityType})
// 			}

// 		wp.Loads = l

// 	return &d, nil
// }

// func DispatchShip(journeyName string) (string, error) {
// 	data, err := os.ReadFile(journeyName + ".json")

// 	if err != nil {
// 		return "", err
// 	}
// 	journey := TravelPlan{}
// 	dataString := string(data)

// 	json.Unmarshal([]byte(dataString), &journey)

// 	detailedJourney := DetailedJourney{}
// 	newRecord := &Trip{}
// 	newRecord.DetailedJourney = detailedJourney
// 	err = mgm.Coll(newRecord).Create(newRecord)
// 	if err != nil {
// 		return "", err
// 	}

// 	return string(newRecord.ID.String()), nil
// }

// func GetJourney(id string) (Trip, error) {
// 	cargo := Trip{}
// 	err := mgm.Coll(&Trip{}).FindByID(id, &cargo)
// 	if err != nil {
// 		return cargo, err
// 	}
// 	return cargo, nil
// }

// func GetAnyJourney() (Trip, error) {
// 	cargo := Trip{}
// 	err := mgm.Coll(&Trip{}).FindOne(mgm.Ctx(), bson.M{}).Decode(&cargo)
// 	if err != nil {
// 		return cargo, err
// 	}
// 	return cargo, nil
// }

// func AnalyzeNeed(need string, dependencies map[string]string) error {
// 	entity, entityType := getEntity(need)
// 	_, ok := dependencies[entity]
// 	if !ok {
// 		dependencies[entity] = entityType
// 	} else {
// 		if dependencies[entity] != "" && entityType != "" && dependencies[entity] != entityType {
// 			return fmt.Errorf("entity type mismatch\n %s already declared with type %s, trying to redeclare as %s", entity, dependencies[entity], entityType)
// 		}
// 		if entityType != "" && dependencies[entity] == "" {
// 			dependencies[entity] = entityType
// 		}
// 	}

// 	return nil
// }

// func AnalyzeTrip(trip *Journey) (map[string]string, error) {
// 	dependencies := map[string]string{}

// 	for _, waypoint := range trip.Waypoints {
// 		for _, container := range waypoint.Loads.Containers {
// 			for _, need := range container.Needs {
// 				err := AnalyzeNeed(need, dependencies)
// 				if err != nil {
// 					return nil, err
// 				}
// 			}

// 		}
// 	}

// 	return dependencies, nil
// }

// func AddCargoToTrip(trip *Trip, name string, payload string) error {
// 	if trip.Cargo.Goods == nil {
// 		trip.Cargo.Goods = map[string]string{}
// 	}
// 	trip.Cargo.Goods[name] = payload
// 	err := mgm.Coll(trip).Update(trip)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

// func ShipMove(trip *Trip) error {
// 	for _, waypoint := range trip.Waypoints {
// 		for _, container := range waypoint.Loads.Containers {
// 			for _, product := range container.Produces {
// 				err := AnalyzeNeed(product, dependencies)
// 				if err != nil {
// 					return err
// 				}
// 			}

// 		}
// 	}
// 	err := AnalyzeTrip(trip)
// 	if err != nil {
// 		return err
// 	}
// 	err = mgm.Coll(trip).Update(trip)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }
