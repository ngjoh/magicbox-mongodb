package journeys

import (
	"encoding/json"
	"testing"

	"github.com/atotto/clipboard"
)

var kitchenName = "nexi-sharepoint"

func TestFindWaypoints(t *testing.T) {

	wayPoints, err := FindWaypoints(kitchenName)
	if err != nil {
		t.Error(err)

	}

	text, _ := json.MarshalIndent(wayPoints, "", "  ")
	clipboard.WriteAll(string(text))

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
	clipboard.WriteAll(string(text))

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
	clipboard.WriteAll(string(text))

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
	clipboard.WriteAll(string(text))

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

	clipboard.WriteAll(mermaid)

}

func TestGetDiagram(t *testing.T) {
	// x
	mermaid, err := GetDiagram(kitchenName, 1)
	if err != nil {
		t.Error(err)
	}
	clipboard.WriteAll(mermaid)

}
