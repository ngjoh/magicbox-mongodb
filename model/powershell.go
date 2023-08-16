package model

import (
	"errors"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/db"
	"github.com/koksmat-com/koksmat/powershell"
	"go.mongodb.org/mongo-driver/bson"
)

type PowershellScript struct {
	mgm.DefaultModel `bson:",inline"`
	Id               string `json:"id"`
	Src              string `json:"src"`
	Description      string `json:"description"`
	Host             string `json:"host"`
	InputSchema      string `json:"inputSchema"`
	OutputSchema     string `json:"outputSchema"`
	ForRoles         string `json:"roles"`
}

func NewPowerShellScript(auth Authorization, id string, src string, description string, inputSchema string, outputSchema string, forRoles string) (*PowershellScript, error) {

	newRecord := &PowershellScript{
		Id:           id,
		DefaultModel: mgm.DefaultModel{},
		Src:          src,
		Description:  description,
		InputSchema:  inputSchema,
		OutputSchema: outputSchema,
		ForRoles:     forRoles,
	}

	err := mgm.Coll(newRecord).Create(newRecord)

	return newRecord, err
}

func GetPowerShellScript(auth Authorization, id string) (*PowershellScript, error) {
	return db.FindOne[*PowershellScript](&PowershellScript{}, bson.D{{"id", id}})
}

func GetPowerShellScripts(auth Authorization) ([]*PowershellScript, error) {
	return db.GetAll[*PowershellScript](&PowershellScript{})
}

func ExecutePowerShellScriptByRef(auth Authorization, id string, params string) (any, error) {
	script, err := GetPowerShellScript(auth, id)
	if err != nil {
		return "", err
	}

	switch script.Host {
	case "exchange":
		return powershell.RunExchange[any](auth.AppId, "scriptid:"+id, params, script.Src, powershell.CallbackMockup)
	case "pnp":
		return powershell.RunPNP[any](auth.AppId, "scriptid:"+id, params, script.Src, powershell.CallbackMockup)

	default:

		return "", errors.New("invalid host")
	}

}

func ExecutePowerShellScript(appId string, host string, script string, params string) (string, error) {

	switch host {
	case "exchange":
		res, err := powershell.RunRawExchange(appId, "adhoc", params, script, powershell.CallbackMockup)
		if err != nil {
			return "", err
		}
		return res, nil
	case "pnp":
		res, err := powershell.RunRawPNP(appId, "adhoc", params, script, powershell.CallbackMockup)
		if err != nil {
			return "", err
		}
		return res, nil
	default:

		return "", errors.New("invalid host")
	}

}
