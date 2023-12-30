package kitchen

import (
	"bytes"
	"os"
	"path"
	"strings"

	"github.com/yuin/goldmark"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func getMetadataProperty(meta map[string]interface{}, key string) string {
	if meta[key] != nil {
		return meta[key].(string)
	}
	return ""
}
func ReadMarkdown(filename string) (string, map[string]interface{}, error) {
	var buf bytes.Buffer
	if !fileExists(filename) {
		return "", nil, nil
	}
	fileContent, err := os.ReadFile(filename)
	if err != nil {
		return "", nil, err
	}
	context := parser.NewContext()
	md := goldmark.New(
		goldmark.WithExtensions(extension.GFM, meta.Meta),
		goldmark.WithParserOptions(
			parser.WithAutoHeadingID(),
		),
		goldmark.WithRendererOptions(
			html.WithHardWraps(),
			html.WithXHTML(),
		),
	)
	if err := md.Convert(fileContent, &buf, parser.WithContext(context)); err != nil {
		return "", nil, err
	}
	metaData := meta.Get(context)
	// title := metaData["Title"]
	return buf.String(), metaData, nil
}
func List() (*[]Kitchen, error) {
	userHome, err := os.UserHomeDir()
	if err != nil {
		return nil, err
	}

	root := path.Join(userHome, "kitchens")
	dirs, err := os.ReadDir(root)
	if err != nil {
		return nil, err
	}

	kitchens := []Kitchen{}
	meta := map[string]interface{}{}
	for _, dir := range dirs {
		if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") {

			kitchen := Kitchen{
				Name: dir.Name(),
				Path: path.Join(root, dir.Name()),
			}
			stations, err := os.ReadDir(path.Join(root, dir.Name()))
			if err != nil {
				return nil, err
			}
			for _, station := range stations {
				if station.IsDir() {
					stationPath := path.Join(root, dir.Name(), station.Name())
					station := Station{
						Name: station.Name(),
						Path: stationPath,
					}

					station.Readme, meta, _ = ReadMarkdown(path.Join(stationPath, "readme.md"))
					station.Description = getMetadataProperty(meta, "description")

					kitchen.Stations = append(kitchen.Stations, station)
				}
			}

			kitchen.Readme, meta, _ = ReadMarkdown(path.Join(root, dir.Name(), "readme.md"))
			kitchen.Description = getMetadataProperty(meta, "description")

			kitchens = append(kitchens, kitchen)
			//	}
		}
	}

	return &kitchens, nil
}
