package io

import (
	"log"
	"os"

	"time"
)

func Waitfor(filePath string) {
	for {
		_, err := os.Open(filePath)

		if err == nil {
			log.Println("Found", filePath)
			log.Println("Waiting 2 secs")
			time.Sleep(2 * time.Second)
			log.Println("Continuing")
			return
		}

		time.Sleep(5 * time.Second)
		log.Println("Waiting for", filePath)
	}

}
