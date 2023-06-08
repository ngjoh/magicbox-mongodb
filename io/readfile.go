package io

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

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
	if len(body) < 1 {
		return recipients
	}
	jsonErr := json.Unmarshal(body, &recipients)

	if jsonErr != nil {
		log.Fatal(jsonErr)
	}
	defer jsonFile.Close()
	return recipients
}

func GetFileNames(pathname string, prefix string) []string {
	fileNames := []string{}
	err := filepath.Walk(pathname, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			fmt.Println(err)
			return err
		}
		fileName := info.Name()
		if strings.HasPrefix(fileName, prefix) {
			_ = append(fileNames, "s")
		}

		return nil
	})
	if err != nil {
		log.Println(err)
	}
	return fileNames
}
