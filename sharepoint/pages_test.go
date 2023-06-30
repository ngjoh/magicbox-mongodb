package sharepoint

import (
	"testing"

	"github.com/spf13/viper"
)

func TestGetIntranetsPages(t *testing.T) {
	_, err := Pages(viper.GetString("PNPSITE"))

	if err != nil {
		t.Error(err)
	}

}
