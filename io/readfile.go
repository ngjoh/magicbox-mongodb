package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func goDotEnvVariable(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Readfile[K any](filePath string) []K {
	jsonFile, err := os.Open(filePath)

	if err != nil {
		fmt.Println(err)
	}
	body, readErr := ioutil.ReadAll(jsonFile)

	if readErr != nil {
		panic(readErr)
	}

	recipients := []K{}
	jsonErr := json.Unmarshal(body, &recipients)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	defer jsonFile.Close()
	return recipients
}
