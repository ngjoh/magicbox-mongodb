package sharepoint

import (
	"testing"
)

func TestSharePoint(t *testing.T) {
	Ping("https://christianiabpos.sharepoint.com/sites/cava3")
	err := FilterTemplate("assets/template.xml")

	if err != nil {
		t.Fatalf("Should not return error")
	}

}
