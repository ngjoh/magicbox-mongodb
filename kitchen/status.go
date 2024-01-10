package kitchen

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type Pair struct {
	Key   string
	Value string
}

type Environment struct {
	Type   string
	Tenant string
	Name   string
	Title  string
	Pairs  []Pair
}

type Status struct {
	Name         string        `json:"name"`
	Title        string        `json:"title"`
	About        string        `json:"about"`
	Description  string        `json:"description"`
	Url          string        `json:"url"`
	Environments []Environment `json:"environments"`
}

func GetStatus(kitchen string) (Status, error) {
	root := viper.GetString("KITCHENROOT")
	status := Status{}
	kitchenPath := path.Join(root, kitchen)
	about, meta, err := ReadMarkdown(kitchenPath, "readme.md")
	if err != nil {
		return status, err
	}
	status.About = about
	status.Title = GetMetadataProperty(meta, "title", kitchen)
	status.Description = GetMetadataProperty(meta, "description", "")
	sharePointPath := path.Join(kitchenPath, ".koksmat", "sharepoint")

	sharePointEnvironments, err := os.ReadDir(sharePointPath)
	if err != nil {
		return status, nil
	}
	for _, c := range sharePointEnvironments {
		if c.IsDir() {
			env := Environment{}
			env.Name = c.Name()
			env.Title = c.Name()
			env.Type = "sharepoint"
			env.Tenant = c.Name()
			env.Pairs = []Pair{}
			status.Environments = append(status.Environments, env)

		}
	}

	return status, nil
}
