package journeys

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

func Pack(filepath string) error {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io

	// find the last element of a split string

	split := strings.Split(filepath, "/")
	filename := split[len(split)-1]

	journeyDir, err := os.ReadDir(filepath)

	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Journey Directory", filename)
	journey := Journey{}
	for _, v := range journeyDir {

		n := strings.Join(strings.Split(v.Name(), " ")[1:], "")
		journey.Waypoints = append(journey.Waypoints, Waypoint{Port: n})
		log.Println(n)

	}

	//

	file, err := json.MarshalIndent(journey, "", " ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(filename+".json", file, 0644)
	return err
}
func Unpack(fileName string) error {
	// Make sure to load your metered License API key prior to using the library.
	// If you need a key, you can sign up and create a free one at https://cloud.unidoc.io
	return nil
}

func Run() {

}
