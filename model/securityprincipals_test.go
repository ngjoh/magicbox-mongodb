package model

import (
	"log"
	"os"
	"testing"

	"github.com/koksmat-com/koksmat/config"
)

func TestHelloName(t *testing.T) {

	key, hash, _ := IssueAccessKey("tester")
	log.Println(key)
	log.Println(hash)
	valid, newHash := Authenticate("tester", key)
	log.Println(newHash)
	if !valid {
		t.Fatalf("Didn't work")
	}
}

func TestMain(m *testing.M) {
	config.Setup("../.env")
	code := m.Run()

	os.Exit(code)
}
