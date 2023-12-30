package kitchen

import (
	"encoding/json"
	"log"
	"testing"
)

func TestRun(t *testing.T) {

	k, err := List()
	if err != nil {
		t.Error(err)
	}
	t.Log(k)
	//text := string("dafdsf")

	j, err := json.Marshal(k)
	log.Println(string(j))

}
