package powershell

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"embed"
)

//go:embed scripts
var scripts embed.FS

func Run(fileName, args string) (result string, consoleText string, err error) {
	cmd := exec.Command("pwsh", "-nologo", "-noprofile")

	dir := ".koksmat/powershell"
	os.MkdirAll(dir, os.ModePerm)
	cmd.Dir = dir

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return "", "", err
	}

	defer stdin.Close()

	log.Println("Connecting to Exchange Online")
	ps1Code, err := scripts.ReadFile("scripts/connectors/exchange-test.ps1")
	if err != nil {

		return "", "", err
	}
	err = os.WriteFile(path.Join(dir, "init.ps1"), ps1Code, 0644)
	if err != nil {
		return "", "", err
	}

	ps2Code, err := scripts.ReadFile(fileName)
	if err != nil {
		return "", "", err
	}
	err = os.WriteFile(path.Join(dir, "run.ps1"), ps2Code, 0644)
	if err != nil {
		return "", "", err
	}
	fmt.Fprintln(stdin, ". ./init.ps1")
	script := fmt.Sprintf(`. ./run.ps1  %s`, args)

	log.Println("Executing", script)
	fmt.Fprintln(stdin, script)

	//out, err := cmd.Output()

	out, err := cmd.CombinedOutput()

	if err != nil {
		return "", "", err
	}

	fmt.Printf("OUTPUT: %s\n", out)
	err = os.WriteFile(path.Join(dir, "output.txt"), out, 0644)
	if err != nil {
		return "", "", err
	}

	outputJson, err := os.ReadFile(path.Join(dir, "output.json"))
	if err != nil {
		return "", "", err
	}
	result = fmt.Sprintf("%s", outputJson)
	return result, fmt.Sprintf("%s", out), nil
}
