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

func Run[R any](fileName, args string) (result *R, err error) {
	cmd := exec.Command("pwsh", "-nologo", "-noprofile")
	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	dir := ".koksmat/powershell"
	os.MkdirAll(dir, os.ModePerm)
	cmd.Dir = dir

	os.Remove(path.Join(dir, "output.json"))
	ps1Code, err := scripts.ReadFile("scripts/connectors/exchange-test.ps1")
	if err != nil {

		return nil, err
	}
	ps2Code, err := scripts.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path.Join(dir, "run.ps1"), ps2Code, 0644)
	if err != nil {
		return nil, err
	}
	err = os.WriteFile(path.Join(dir, "init.ps1"), ps1Code, 0644)
	if err != nil {
		return nil, err
	}
	script := fmt.Sprintf(`. ./run.ps1  %s`, args)
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, ". ./init.ps1")
		fmt.Fprintln(stdin, script)

	}()
	_, err = cmd.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}
	dataOut := new(R)
	outputJson, err := os.ReadFile(path.Join(dir, "output.json"))
	if err == nil {

		jsonErr := json.Unmarshal(outputJson, &dataOut)
		if jsonErr != nil {
			s := fmt.Sprintf("[%s]", outputJson)
			outArray := []byte(s)
			jsonErr := json.Unmarshal(outArray, &dataOut)
			if jsonErr != nil {
				return nil, jsonErr
			}
		}
	}

	result = *&dataOut // fmt.Sprintf("%s", outputJson)
	return result, nil
}
