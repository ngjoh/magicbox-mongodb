package audit

import (
	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/config"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const AuditlogCollectionName = "audit_log"

type AuditLog struct {
	mgm.DefaultModel `bson:",inline"`
	Database         string `json:"database"`
	AppId            string `json:"appid"`
	Subject          string `json:"subject"`
}

type PowerShellLog struct {
	mgm.DefaultModel `bson:",inline"`
	Database         string `json:"database"`
	AppId            string `json:"appid"`
	Subject          string `json:"subject"`
	ScriptName       string `json:"scriptname"`
	ScriptSrc        string `json:"scriptsrc"`
	Input            string `json:"input"`
	Output           string `json:"result"`
	HasError         bool   `json:"haserror"`
	Console          string `json:"console"`
}

// func (model *PowerShellLog) CollectionName() string {
// 	return AuditlogCollectionName

// }
func (auditLog *PowerShellLog) Collection() *mgm.Collection {
	// Create new client

	client, err := mgm.NewClient(options.Client().ApplyURI(config.MongoConnectionString()))
	if err != nil {
		panic(err)
	}
	// Get the model's db
	db := client.Database("magicbox")
	// return the model's custom collection
	return mgm.NewCollection(db, "audit_log")
}
func (auditLog *AuditLog) Collection() *mgm.Collection {
	// Create new client

	client, err := mgm.NewClient(options.Client().ApplyURI(config.MongoConnectionString()))
	if err != nil {
		panic(err)
	}
	// Get the model's db
	db := client.Database("magicbox")
	// return the model's custom collection
	return mgm.NewCollection(db, "audit_log")
}

func LogAudit(app string, subject string) {

	newRecord := &AuditLog{
		Database: config.DatabaseName(),
		AppId:    app,
		Subject:  subject,
	}

	mgm.Coll(newRecord).Create(newRecord)

}

func LogPowerShell(app string, scriptName string, scriptSrc string, input string, output string, hasError bool, console string) {

	newRecord := &PowerShellLog{
		DefaultModel: mgm.DefaultModel{},
		Database:     config.DatabaseName(),
		AppId:        app,
		Subject:      "powershell",
		ScriptName:   scriptName,
		ScriptSrc:    scriptSrc,
		Input:        input,
		Output:       output,
		HasError:     hasError,
		Console:      console,
	}

	err := mgm.Coll(newRecord).Create(newRecord)
	if err != nil {
		panic(err)
	}

}
