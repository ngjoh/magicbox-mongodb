package kitchen

import (
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type Meta struct {
	Title       string                `json:"title"`
	Description string                `json:"description"`
	Environment []string              `json:"environment"`
	Parameters  []PowershellParameter `json:"parameters"`
	Connections []string              `json:"connections"`
}

func GetMetadata(kitchenName string, stationName string, filename string) (Meta, error) {
	root := viper.GetString("KITCHENROOT")

	file := path.Join(root, kitchenName, stationName, filename)

	markdown, environmentVariablesReferencedInScript, err := ReadMarkdownFromPowerShell(file)
	if err != nil {
		fmt.Println(err)
	}
	_, meta, err := ParseMarkdown(filepath.Dir(file), markdown)
	if err != nil {
		fmt.Println(err)
	}
	parameters, err := GetPowerShellFileParameters(file)
	if err != nil {
		fmt.Println(err)
	}

	metadata := Meta{
		Title:       GetMetadataProperty(meta, "title", ""),
		Description: GetMetadataProperty(meta, "description", ""),
		Environment: environmentVariablesReferencedInScript,
		Parameters:  parameters,
	}

	connections := GetMetadataProperty(meta, "connection", "")
	if connections != "" {
		metadata.Connections = strings.Split(connections, ",")
	}
	return metadata, nil
}
