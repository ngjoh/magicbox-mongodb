package connectors

import (
	"encoding/json"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type Current struct {
	Name       string `json:"name"`
	SharePoint string `json:"sharepoint"`
	Path       string `json:"path"`
	GitOrg     string `json:"gitorg"`
	Mongo      string `json:"mongo"`
}

type Sharepoint struct {
	Tenant string `json:"tenant"`
	Site   string `json:"site"`
}
type Mongo struct {
	Cluster  string `json:"cluster"`
	Database string `json:"database"`
}
type MateContext struct {
	Current    Current      `json:"current"`
	SharePoint []Sharepoint `json:"sharepoint"`
	Mongo      []Mongo      `json:"mongo"`
}

func GetMateContext() (*MateContext, error) {
	kitchenRoot := viper.GetString("KITCHENROOT")
	mateContextPath := path.Join(kitchenRoot, "mate.json")
	bytes, err := os.ReadFile(mateContextPath)
	if err != nil {
		return nil, err
	}

	mateContext := MateContext{}

	err = json.Unmarshal(bytes, &mateContext)
	if err != nil {
		return nil, err
	}
	mateContext.Current.Path = path.Join(mateContextPath, mateContext.Current.Name)
	return &mateContext, nil
}
func SharePointTenants() ([]Connector, error) {

	kitchenRoot := viper.GetString("KITCHENROOT")

	mateContext, err := GetMateContext()
	if err != nil {
		return nil, err
	}

	sharePointConnectorPath := path.Join(kitchenRoot, mateContext.Current.Name, ".sharepoint", "tenants")
	dirs, err := os.ReadDir(sharePointConnectorPath)
	if err != nil {
		return nil, err
	}
	connectors := make([]Connector, 0)
	for _, dir := range dirs {
		if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") {
			connector := Connector{
				Name:      dir.Name(),
				Url:       "file://" + path.Join(sharePointConnectorPath, dir.Name()),
				IsCurrent: dir.Name() == mateContext.Current.SharePoint,
			}
			connectors = append(connectors, connector)
		}
	}
	return connectors, nil
}
