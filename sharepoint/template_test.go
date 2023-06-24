package sharepoint

import (
	"testing"
)

func TestFilterTemplate(t *testing.T) {

	err := FilterTemplate("assets/template.xml")

	if err != nil {
		t.Fatalf("Should not return error")
	}

}

func TestFilterTemplate2(t *testing.T) {

	err := FilterTemplate("assets/template-filtered.xml")

	if err != nil {
		t.Fatalf("Should not return error")
	}

}
