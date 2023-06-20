package powershell

import (
	"embed"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"path"

	"github.com/koksmat-com/koksmat/audit"
)

//go:embed scripts
var scripts embed.FS

func Execute(fileName, args string) (output []byte, err error, console string) {
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

		return nil, err, ""
	}
	ps2Code, err := scripts.ReadFile(fileName)
	if err != nil {
		return nil, err, ""
	}
	err = os.WriteFile(path.Join(dir, "run.ps1"), ps2Code, 0644)
	if err != nil {
		return nil, err, ""
	}
	err = os.WriteFile(path.Join(dir, "init.ps1"), ps1Code, 0644)
	if err != nil {
		return nil, err, ""
	}
	script := fmt.Sprintf(`. ./run.ps1  %s`, args)
	go func() {
		defer stdin.Close()
		fmt.Fprintln(stdin, ". ./init.ps1")
		fmt.Fprintln(stdin, script)

	}()
	srcCode := fmt.Sprintf("[%s]", ps2Code)
	combinedOutput, err := cmd.CombinedOutput()
	if err != nil {
		audit.LogPowerShell("koksmat", fileName, srcCode, args, "", true, string(combinedOutput))
		log.Println(fmt.Sprint(err) + ": " + string(combinedOutput))
		return nil, errors.New("Could not run PowerShell script"), string(combinedOutput)
	}

	outputJson, err := os.ReadFile(path.Join(dir, "output.json"))

	audit.LogPowerShell("koksmat", fileName, srcCode, args, fmt.Sprintf("[%s]", outputJson), false, string(combinedOutput))

	return outputJson, nil, string(combinedOutput)
}

func Run[R any](fileName string, args string) (result *R, err error) {

	output, err, _ := Execute(fileName, args)
	dataOut := new(R)
	textOutput := fmt.Sprintf("%s", output)
	if (output != nil) && (textOutput != "") {

		jsonErr := json.Unmarshal(output, &dataOut)
		if jsonErr != nil {
			s := fmt.Sprintf("[%s]", output)
			outArray := []byte(s)
			jsonErr := json.Unmarshal(outArray, &dataOut)
			if jsonErr != nil {
				log.Println("Error parsing output: ", jsonErr)
			}
		}
	}
	result = *&dataOut // fmt.Sprintf("%s", outputJson)
	return result, err
}
