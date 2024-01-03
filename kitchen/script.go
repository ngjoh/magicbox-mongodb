package kitchen

import (
	"os"
	"strings"
)

// /Users/nielsgregersjohansen/kitchens/sharepoint-branding/install/20 apply-sitetemplate.ps1

func ParsePowerShellFile(fileContent string) (string, error) {
	markdown := ""
	commentBlock := strings.Split(string(fileContent), "<#")

	for _, block := range commentBlock {
		if strings.Index(block, "#>") == -1 {
			if block != "" {
				markdown += "\n```powershell\n"
				markdown += block
				markdown += "\n```\n"
			}
		} else {
			comments := strings.Split(block, "#>")
			markdown += comments[0]
			if len(comments) > 1 {

				markdown += "\n```powershell\n"
				markdown += comments[1]
				markdown += "\n```\n"

			}
		}

	}
	markdown += ` 
Note that Koksmat will append a few lines of code here which will take the value of $result, convert it to JSON and store it in a file called output.json	
	`
	return markdown, nil
}
func ReadMarkdownFromPowerShell(filepath string) (string, error) {

	if !fileExists(filepath) {
		return "", nil
	}
	fileContent, err := os.ReadFile(filepath)
	if err != nil {
		return "", err
	}
	return ParsePowerShellFile(string(fileContent))
}
