package journeys

import (
	"encoding/json"
	"fmt"
	"testing"
)

var kitchenName = "nexi-sharepoint"

func TestFindWaypoints(t *testing.T) {

	wayPoints, err := FindWaypoints(kitchenName)
	if err != nil {
		t.Error(err)

	}

	text, _ := json.MarshalIndent(wayPoints, "", "  ")
	fmt.Printf("%s\n", text)

}

func TestFindJourneys(t *testing.T) {
	wayPoints, err := FindWaypoints(kitchenName)
	if err != nil {
		t.Error(err)

	}
	journeys, _, err := FindJourneys(*wayPoints)
	if err != nil {
		t.Error(err)

	}

	text, _ := json.MarshalIndent(journeys, "", "  ")
	fmt.Printf("%s\n", text)

}

func TestBuildTravelPlan(t *testing.T) {
	wayPoints, err := FindWaypoints(kitchenName)
	if err != nil {
		t.Error(err)

	}
	journeys, _, err := FindJourneys(*wayPoints)
	if err != nil {
		t.Error(err)

	}

	travelPlan, err := BuildTravelPlan(*journeys)
	if err != nil {
		t.Error(err)
	}

	text, _ := json.MarshalIndent(travelPlan, "", "  ")
	fmt.Printf("%s\n", text)

}

func TestBuildTriggerPlan(t *testing.T) {
	wayPoints, err := FindWaypoints(kitchenName)
	if err != nil {
		t.Error(err)

	}
	_, triggers, err := FindJourneys(*wayPoints)
	if err != nil {
		t.Error(err)

	}

	text, _ := json.MarshalIndent(triggers, "", "  ")
	fmt.Printf("%s\n", text)

}
func TestBuildTravelPlanMermaid(t *testing.T) {
	wayPoints, err := FindWaypoints(kitchenName)
	if err != nil {
		t.Error(err)

	}
	journeys, triggers, err := FindJourneys(*wayPoints)
	if err != nil {
		t.Error(err)

	}

	travelPlan, err := BuildTravelPlan(*journeys)
	if err != nil {
		t.Error(err)
	}

	mermaid, err := BuildTravelPlanMermaid(travelPlan, *triggers)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("%s\n", mermaid)

}

func TestGetDiagram(t *testing.T) {
	// x
	mermaid, err := GetDiagram(kitchenName, 1)
	if err != nil {
		t.Error(err)
	}
	fmt.Printf("%s\n", mermaid)

}
