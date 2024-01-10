package kitchen

import (
	"encoding/json"
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/spf13/viper"
)

type Folder struct {
	Path string `json:"path"`
}
type VSCodeWorkspace struct {
	Folders []Folder `json:"folders"`
}

func CreateKitchen(name string) error {

	root := viper.GetString("KITCHENROOT")
	kitchenPath := path.Join(root, name)

	err := os.MkdirAll(kitchenPath, 0755)
	if err != nil {
		return err
	}

	readmeFilePath := path.Join(kitchenPath, "readme.md")

	if fileExists(readmeFilePath) {
		return fmt.Errorf("kitchen %s already exists", name)
	}
	data := fmt.Sprintf(`---
title: %s
description: Describe the main purpose of this kitchen
---

# %s
`, name, name)

	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(readmeFilePath, []byte(data), 0644)
	if err != nil {
		return err
	}
	return CreateWorkspaceFile()
}

func CreateWorkspaceFile() error {
	root := viper.GetString("KITCHENROOT")
	wsPath := path.Join(root, "kitchens.code-workspace")
	workspaces := VSCodeWorkspace{}
	dirs, err := os.ReadDir(root)
	if err != nil {
		return err
	}

	for _, dir := range dirs {
		if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") {
			n := dir.Name()
			workspaces.Folders = append(workspaces.Folders, Folder{Path: n})
		}
	}

	data, err := json.MarshalIndent(workspaces, "", " ")
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = os.WriteFile(wsPath, data, 0644)
	if err != nil {
		return err
	}
	return nil
}
