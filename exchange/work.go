package exchange

import (
	"log"
	"time"
)

/*
Function designed to never return
*/
func Work() {
	// log.Println("Starting exchange worker")
	for {
		//TODO implement something to do
		time.Sleep(5 * time.Second)
		log.Println("Waiting for jobs")
	}
}
