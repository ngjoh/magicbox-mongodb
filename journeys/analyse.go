package journeys

import (
	"fmt"
	"strings"

	"github.com/koksmat-com/koksmat/kitchen"
)

type ContainerIO struct {
	ContainerName string `json:"containername"`
	Input         string `json:"input"`
	Output        string `json:"output"`
	Trigger       string `json:"trigger"`
	Connection    string `json:"connection"`

	Tag string `json:"tag"`
}

type PortIO struct {
	Portname    string        `json:"portname"`
	Description string        `json:"description"`
	Tag         string        `json:"tag"`
	Containers  []ContainerIO `json:"containers"`
}

type KitchenWaypointsReport struct {
	JourneyName string   `json:"journeyname"`
	Description string   `json:"description"`
	Tag         string   `json:"tag"`
	Ports       []PortIO `json:"ports"`
}

type JourneyReport struct {
	Tag       string `json:"tag"`
	Artifact  string `json:"artifact"`
	Direction string `json:"direction"`
}

type Handover struct {
	From     string `json:"from"`
	To       string `json:"to"`
	Artifact string `json:"artifact"`
}

type ConsumentInfo struct {
	Tag      string `json:"tag"`
	Artifact string `json:"artifact"`
}

type ShipperInfo struct {
	Name      string `json:"name"`
	Order     int    `json:"order"`
	Container string `json:"container"`
	Artifact  string `json:"artifact"`
}

func FindWaypoints(kitchenName string) (*KitchenWaypointsReport, error) {

	_, err := kitchen.GetPath(kitchenName)
	if err != nil {
		return nil, err
	}

	k, err := kitchen.GetStations(kitchenName)
	if err != nil {
		return nil, err
	}

	result := &KitchenWaypointsReport{
		JourneyName: k.Name,
		Description: k.Description,
		Tag:         k.Tag,
	}
	for _, station := range k.Stations {
		fmt.Println(station.Name)
		scripts, err := kitchen.GetScripts(station.Path, "")
		if err != nil {
			return nil, err
		}

		port := PortIO{
			Portname:    station.Name,
			Description: station.Description,
			Tag:         station.Tag,
		}

		for _, script := range scripts {
			container := ContainerIO{
				ContainerName: script.Name,
				Input:         script.Input,
				Output:        script.Output,
				Trigger:       script.Trigger,
				Connection:    script.Connection,
				Tag:           script.Tag,
			}
			port.Containers = append(port.Containers, container)

		}
		result.Ports = append(result.Ports, port)

	}

	return result, nil
}

func FindInput(name string, wayPoints KitchenWaypointsReport) ([]ConsumentInfo, error) {
	result := []ConsumentInfo{}
	if (name == "") || (name == "null") {
		return result, nil
	}
	for _, port := range wayPoints.Ports {
		for _, container := range port.Containers {
			if container.Input == name {

				result = append(result, ConsumentInfo{
					Tag:      fmt.Sprintf("%s.%s", container.Tag, port.Tag),
					Artifact: container.Input,
				})
			}

		}
	}
	return result, nil
}

func FindOutput(name string, wayPoints KitchenWaypointsReport, reportMode bool) ([]string, error) {
	result := []string{}
	if (name == "") || (name == "null") {
		return result, nil
	}
	for _, port := range wayPoints.Ports {
		for _, container := range port.Containers {
			if container.Output == name {
				if !reportMode {
					if len(result) > 1 {
						return nil, fmt.Errorf("Multiple producers for %s - %s.%s", container.Input, container.ContainerName, port.Portname)
					}
				}

				result = append(result, fmt.Sprintf("%s.%s", container.Tag, port.Tag))
			}

		}
	}
	return result, nil
}

func FindShipper(wayPoints KitchenWaypointsReport, artifact string) (string, string) {
	artifactName := ""
	containerName := ""
	for _, port := range wayPoints.Ports {

		for _, container := range port.Containers {
			if container.Output == artifact {
				containerName = fmt.Sprintf("%s.%s", container.Tag, port.Tag)
				artifactName = artifact
			}
		}
	}

	return artifactName, containerName
}

func FindJourneys(wayPoints KitchenWaypointsReport) (*[]JourneyReport, *[]ShipperInfo, error) {

	result := []JourneyReport{}
	triggers := []ShipperInfo{}

	for _, port := range wayPoints.Ports {
		for _, container := range port.Containers {

			inputs := strings.Split(container.Input, ",")
			for _, input := range inputs {
				i := strings.TrimSpace(input)
				if (i != "") && (i != "null") {

					producers, err := FindOutput(i, wayPoints, false)

					if err != nil {
						return nil, nil, err
					}

					if len(producers) == 1 {

						result = append(result, JourneyReport{
							//Container:   container,
							Tag: producers[0],

							Artifact:  i,
							Direction: "out",
						})

					}
				}
			}
		}
	}

	for _, port := range wayPoints.Ports {
		for _, container := range port.Containers {

			consumers, err := FindInput(container.Input, wayPoints)

			if err != nil {
				return nil, nil, err
			}
			for _, consumer := range consumers {

				result = append(result, JourneyReport{
					//Container:   container,

					Tag:       consumer.Tag,
					Artifact:  consumer.Artifact,
					Direction: "in",
				})

			}
		}
	}
	for _, port := range wayPoints.Ports {
		for _, container := range port.Containers {
			if (container.Trigger != "") && (container.Trigger != "null") {

				ts := strings.Split(container.Trigger, ",")
				for ix, trigger := range ts {
					t := strings.TrimSpace(trigger)
					_, shipper := FindShipper(wayPoints, t)
					if (shipper != "") && (shipper != "null") {
						triggers = append(triggers, ShipperInfo{
							Order:     ix + 1,
							Container: shipper,
							Artifact:  t,
							Name:      container.Tag,
						})
					}

				}

			}
		}
	}
	return &result, &triggers, nil
}

func BuildTravelPlan(journey []JourneyReport) ([]Handover, error) {
	result := []Handover{}
	for _, j := range journey {
		if j.Direction == "out" {
			for _, k := range journey {
				if k.Direction == "in" && j.Artifact == k.Artifact {
					result = append(result, Handover{

						From:     j.Tag,
						To:       k.Tag,
						Artifact: j.Artifact,
					})
				}
			}
		}
	}
	// for _, j := range journey {
	// 	if j.Direction == "in" {
	// 		for _, k := range journey {
	// 			if k.Direction == "out" && j.Artifact == k.Artifact {
	// 				result = append(result, Handover{
	// 					To:       j.Tag,
	// 					From:     k.Tag,
	// 					Artifact: j.Artifact,
	// 				})
	// 			}
	// 		}
	// 	}
	// }
	return result, nil

}

func BuildTriggers(journey []JourneyReport) ([]Handover, error) {
	result := []Handover{}
	for _, j := range journey {
		if j.Direction == "out" {
			for _, k := range journey {
				if k.Direction == "in" && j.Artifact == k.Artifact {
					result = append(result, Handover{

						From:     j.Tag,
						To:       k.Tag,
						Artifact: j.Artifact,
					})
				}
			}
		}
	}
	// for _, j := range journey {
	// 	if j.Direction == "in" {
	// 		for _, k := range journey {
	// 			if k.Direction == "out" && j.Artifact == k.Artifact {
	// 				result = append(result, Handover{
	// 					To:       j.Tag,
	// 					From:     k.Tag,
	// 					Artifact: j.Artifact,
	// 				})
	// 			}
	// 		}
	// 	}
	// }
	return result, nil

}

func BuildTravelPlanMermaid(travelPlan []Handover, triggers []ShipperInfo) (string, error) {
	result := "graph TD\n"
	for _, t := range travelPlan {
		result += fmt.Sprintf("%s --> |%s| %s \n", t.From, t.Artifact, t.To)
	}
	for _, t := range triggers {
		result += fmt.Sprintf("%s((%s)) -->  |%d| %s \n", t.Name, t.Name, t.Order, t.Container)
	}
	return result, nil
}
func GetDiagram(kitchen string, version int) (string, error) {
	wayPoints, err := FindWaypoints(kitchen)
	if err != nil {
		return "", err

	}
	journeys, triggers, err := FindJourneys(*wayPoints)
	if err != nil {
		return "", err

	}
	travelPlan, err := BuildTravelPlan(*journeys)
	if err != nil {
		return "", err

	}
	mermaid, err := BuildTravelPlanMermaid(travelPlan, *triggers)
	if err != nil {
		return "", err

	}
	return mermaid, nil

}
