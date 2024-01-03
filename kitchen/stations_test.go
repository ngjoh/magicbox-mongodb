package kitchen

import (
	"log"
	"testing"
)

func TestSpawn(t *testing.T) {

	dir, pid, err := StartSession("sharepoint-branding", "install")
	if err != nil {
		t.Error(err)
	}
	log.Println(dir, pid)
	t.Log(dir, pid)

}
