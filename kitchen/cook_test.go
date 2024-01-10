package kitchen

import (
	"path"
	"testing"

	"github.com/spf13/viper"
)

func TestSetupSession(t *testing.T) {
	root := viper.GetString("KITCHENROOT")
	kitchenPath := path.Join(root, "sharepoint-branding")
	s, err := SetupSessionPath(kitchenPath)
	if err != nil {
		t.Error(err)
	}
	t.Log(s)
}

func TestGetEnvironment(t *testing.T) {
	root := viper.GetString("KITCHENROOT")
	kitchenPath := path.Join(root, "sharepoint-branding")

	s, err := getEnvironmentFilePath(kitchenPath, "365adm")
	if err != nil {
		t.Error(err)
	}
	t.Log(s)

	e := PowerShellEnvironmentVariables(s)
	t.Log(e)
}

// func TestCook(t *testing.T) {
// 	root := viper.GetString("KITCHENROOT")

// 	result, err := Cook("", root, `

// $result = "Hello World"

// start-sleep -s 1

// 	`, nil)
// 	if err != nil {
// 		t.Error(err)

// 	}
// 	t.Log(result)
// }
