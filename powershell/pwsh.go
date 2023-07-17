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
	"github.com/koksmat-com/koksmat/sharepoint"
	"github.com/spf13/viper"
)

//go:embed scripts
var scripts embed.FS

type Setup func() (string, []string, error)

func PwshCwd() string {
	dir := ".koksmat/powershell"
	os.MkdirAll(dir, os.ModePerm)
	return dir
}

func Execute(appId string, fileName, args string, setEnvironment Setup) (output []byte, err error, console string,
) {
	cmd := exec.Command("pwsh", "-nologo", "-noprofile")

	initScript, environment, err := setEnvironment()
	if err != nil {
		return nil, err, ""
	}

	cmd.Env = environment

	stdin, err := cmd.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	cmd.Dir = PwshCwd()

	os.Remove(path.Join(cmd.Dir, "output.json"))
	ps1Code, err := scripts.ReadFile(fmt.Sprintf("scripts/connectors/%s.ps1", initScript))
	if err != nil {

		return nil, err, ""
	}
	ps2Code, err := scripts.ReadFile(fileName)
	if err != nil {
		return nil, err, ""
	}
	err = os.WriteFile(path.Join(cmd.Dir, "run.ps1"), ps2Code, 0644)
	if err != nil {
		return nil, err, ""
	}
	err = os.WriteFile(path.Join(cmd.Dir, "init.ps1"), ps1Code, 0644)
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
		audit.LogPowerShell(appId, fileName, srcCode, args, "", true, string(combinedOutput))
		log.Println(fmt.Sprint(err) + ": " + string(combinedOutput))
		return nil, errors.New("Could not run PowerShell script"), string(combinedOutput)
	}

	outputJson, err := os.ReadFile(path.Join(cmd.Dir, "output.json"))

	audit.LogPowerShell(appId, fileName, srcCode, args, fmt.Sprintf("[%s]", outputJson), false, string(combinedOutput))

	return outputJson, nil, string(combinedOutput)
}

func Run[R any](appId string, fileName string, args string, setup Setup) (result *R, err error) {

	output, err, _ := Execute(appId, fileName, args, setup)
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

var SetupExchange = func() (string, []string, error) {
	env := os.Environ()

	env = append(env, fmt.Sprintf("EXCHCERTIFICATEPASSWORD=%s", viper.GetString("EXCHCERTIFICATEPASSWORD")))
	env = append(env, fmt.Sprintf("EXCHAPPID=%s", viper.GetString("EXCHAPPID")))
	env = append(env, fmt.Sprintf("EXCHORGANIZATION=%s", viper.GetString("EXCHORGANIZATION")))
	env = append(env, fmt.Sprintf("EXCHCERTIFICATE=%s", viper.GetString("EXCHCERTIFICATE")))
	return "exchange", env, nil

}

var SetupPNP = func() (string, []string, error) {

	ps2Code, err := sharepoint.Assets.ReadFile("assets/template-filtered.xml")
	if err != nil {
		return "", []string{}, err
	}
	err = os.WriteFile(path.Join(PwshCwd(), "template.xml"), ps2Code, 0644)
	if err != nil {
		return "", []string{}, err
	}
	env := os.Environ()

	env = append(env, fmt.Sprintf("PNPAPPID=%s", viper.GetString("PNPAPPID")))
	env = append(env, fmt.Sprintf("PNPTENANTID=%s", viper.GetString("PNPTENANTID")))
	env = append(env, fmt.Sprintf("PNPSITE=%s", viper.GetString("PNPSITE")))
	env = append(env, fmt.Sprintf("PNPCERTIFICATE=%s", viper.GetString("PNPCERTIFICATE")))
	return "pnp", env, nil

}

func RunExchange[R any](appId string, fileName string, args string) (result *R, err error) {

	return Run[R](appId, fileName, args, SetupExchange)
}

func RunPNP[R any](appId string, fileName string, args string) (result *R, err error) {

	return Run[R](appId, fileName, args, SetupPNP)
}
