package powershell

import (
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"
)

//go:embed scripts
var scripts embed.FS

func Run2(fileName, args string) (result string, consoleText string, err error) {
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
	fmt.Fprintln(stdin, "write-host '---'")
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

func Run[R any](fileName, args string) (result *R, consoleText string, err error) {
	cmd := exec.Command("pwsh", "-nologo", "-noprofile")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	dir := ".koksmat/powershell"
	os.MkdirAll(dir, os.ModePerm)
	cmd.Dir = dir

	ps1Code, err := scripts.ReadFile("scripts/connectors/exchange-test.ps1")
	if err != nil {

		return nil, "", err
	}
	ps2Code, err := scripts.ReadFile(fileName)
	if err != nil {
		return nil, "", err
	}
	err = os.WriteFile(path.Join(dir, "run.ps1"), ps2Code, 0644)
	if err != nil {
		return nil, "", err
	}
	err = os.WriteFile(path.Join(dir, "init.ps1"), ps1Code, 0644)
	if err != nil {
		return nil, "", err
	}
	script := fmt.Sprintf(`. ./run.ps1  %s`, args)
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, ". ./init.ps1")
		fmt.Fprintln(stdin, script)

	}()
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	outputJson, err := os.ReadFile(path.Join(dir, "output.json"))
	if err != nil {
		return nil, "", err
	}
	dataOut := new(R)
	jsonErr := json.Unmarshal(outputJson, &dataOut)
	if jsonErr != nil {
		return nil, "", jsonErr
	}
	result = *&dataOut // fmt.Sprintf("%s", outputJson)
	return result, fmt.Sprintf("%s", out), nil
}
