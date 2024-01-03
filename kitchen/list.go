package kitchen

import (
	"bytes"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
	img64 "github.com/tenkoh/goldmark-img64"
	"github.com/yuin/goldmark"
	highlighting "github.com/yuin/goldmark-highlighting"
	meta "github.com/yuin/goldmark-meta"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/renderer/html"
)

func GetMetadataProperty(meta map[string]interface{}, key string, defaultValue string) string {

	if meta[key] != nil {
		return meta[key].(string)
	}
	return defaultValue
}

type Metadata map[string]interface{}

func ParseMarkdown(content string) (string, Metadata, error) {
	var buf bytes.Buffer
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
		goldmark.WithExtensions(
			highlighting.Highlighting,
		),
		goldmark.WithExtensions(img64.Img64),
	)
	if err := md.Convert([]byte(content), &buf, parser.WithContext(context)); err != nil {
		return "", nil, err
	}
	metaData := meta.Get(context)
	return buf.String(), metaData, nil
}
func ReadMarkdown(pathname string, filename string) (string, Metadata, error) {

	filepath := filepath.Join(pathname, filename)
	if !fileExists(filepath) {
		return "", nil, nil
	}
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return "", nil, err
	}

	return ParseMarkdown(string(fileContent))

}

func List() (*[]Kitchen, error) {
	// userHome, err := os.UserHomeDir()
	// if err != nil {
	// 	return nil, err
	// }

	// root := path.Join(userHome, "kitchens")
	root := viper.GetString("KITCHENROOT")
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
			// stations, err := os.ReadDir(path.Join(root, dir.Name()))
			// if err != nil {
			// 	return nil, err
			// }
			// for _, station := range stations {
			// 	if station.IsDir() {
			// 		stationPath := path.Join(root, dir.Name(), station.Name())
			// 		station := Station{
			// 			Name: station.Name(),
			// 			Path: stationPath,
			// 		}

			// 		station.Readme, meta, _ = ReadMarkdown(path.Join(stationPath, "readme.md"))
			// 		station.Description = getMetadataProperty(meta, "description")

			// 		kitchen.Stations = append(kitchen.Stations, station)
			// 	}
			// }

			kitchen.Readme, meta, _ = ReadMarkdown(path.Join(root, dir.Name()), "readme.md")
			kitchen.Description = GetMetadataProperty(meta, "description", "")

			kitchens = append(kitchens, kitchen)
			//	}
		}
	}

	return &kitchens, nil
}

func GetScripts(stationPath string, subPath string) ([]Script, error) {

	filePath := path.Join(stationPath, subPath)

	scripts, err := os.ReadDir(filePath)
	result := []Script{}
	if err != nil {
		return nil, err
	}
	for _, script := range scripts {
		if script.IsDir() {
			s, err := GetScripts(stationPath, path.Join(subPath, script.Name()))
			if err != nil {
				return nil, err
			}
			result = append(result, s...)
		}
		if !script.IsDir() && strings.HasSuffix(script.Name(), ".ps1") && !strings.HasPrefix(script.Name(), "test-") {
			markdown, err := ReadMarkdownFromPowerShell(path.Join(filePath, script.Name()))
			if err != nil {
				return nil, err
			}

			_, scriptmeta, _ := ParseMarkdown(markdown)

			script := Script{
				Name:        path.Join(subPath, script.Name()),
				Title:       path.Join(subPath, GetMetadataProperty(scriptmeta, "title", script.Name())),
				Description: GetMetadataProperty(scriptmeta, "description", ""),
			}
			result = append(result, script)

		}
	}
	return result, nil
}

func GetStations(kitchenName string) (*Kitchen, error) {

	root := viper.GetString("KITCHENROOT")
	filepath := filepath.Join(root, kitchenName)
	dirs, err := os.ReadDir(filepath)
	if err != nil {
		return nil, err
	}

	kitchen := &Kitchen{
		Name: kitchenName,
		Path: filepath,
	}

	readme, meta, _ := ReadMarkdown(filepath, "readme.md")
	kitchen.Readme = readme
	kitchen.Title = GetMetadataProperty(meta, "title", kitchenName)
	kitchen.Description = GetMetadataProperty(meta, "description", "")

	for _, dir := range dirs {
		if dir.IsDir() && !strings.HasPrefix(dir.Name(), ".") {
			stationPath := path.Join(filepath, dir.Name())
			stationReadme, statoinmeta, _ := ReadMarkdown(path.Join(stationPath), "readme.md")

			station := Station{
				Readme: stationReadme,
				Name:   dir.Name(),
				Path:   stationPath,
				Title:  GetMetadataProperty(statoinmeta, "title", dir.Name()),
			}

			station.Description = GetMetadataProperty(meta, "description", "")
			station.Scripts, err = GetScripts(stationPath, "")
			/*
				scripts, err := os.ReadDir(stationPath)
				if err != nil {
					return nil, err
				}
				for _, script := range scripts {
					if !script.IsDir() && strings.HasSuffix(script.Name(), ".ps1") && !strings.HasPrefix(script.Name(), "test-") {
						markdown, err := ReadMarkdownFromPowerShell(path.Join(stationPath, script.Name()))
						if err != nil {
							return nil, err
						}

						_, scriptmeta, _ := ParseMarkdown(markdown)

						script := Script{
							Name:        script.Name(),
							Title:       getMetadataProperty(scriptmeta, "title", script.Name()),
							Description: getMetadataProperty(scriptmeta, "description", ""),
						}
						station.Scripts = append(station.Scripts, script)

					}
				}
			*/
			kitchen.Stations = append(kitchen.Stations, station)
			//	}
		}
	}

	return kitchen, nil
}
