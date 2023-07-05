package audit

import (
	"context"

	"github.com/kamva/mgm/v3"
	"github.com/koksmat-com/koksmat/config"
	"github.com/koksmat-com/koksmat/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const AuditlogCollectionName = "audit_log"

type AuditLog struct {
	mgm.DefaultModel `bson:",inline"`
	Database         string `json:"database"`
	AppId            string `json:"appid"`
	Subject          string `json:"subject"`
}

type AuditLogSum struct {
	mgm.DefaultModel `bson:",inline"`
	Date             string `json:"date"`
	Subject          string `json:"subject"`
	Hour             string `json:"hour"`
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

func GetAuditLogs(day string) ([]*AuditLogSum, error) {
	return db.GetFiltered[*AuditLogSum](&AuditLogSum{}, bson.D{{"date", day}})
}

func GetPowerShellLog(id string) (*PowerShellLog, error) {
	return db.FindOneById[*PowerShellLog](&PowerShellLog{}, id)
}

func Aggregate() error {
	ctx := context.TODO()
	pipeline := bson.A{
		bson.D{
			{"$addFields",
				bson.D{
					{"datepart",
						bson.D{
							{"$substr",
								bson.A{
									"$created_at",
									0,
									10,
								},
							},
						},
					},
					{"hour",
						bson.D{
							{"$substr",
								bson.A{
									"$created_at",
									11,
									2,
								},
							},
						},
					},
				},
			},
		},
		bson.D{
			{"$group",
				bson.D{
					{"_id", "$datepart"},
					{"subject", bson.D{{"$first", "$subject"}}},
					{"hour", bson.D{{"$first", "$hour"}}},
					{"count", bson.D{{"$sum", 1}}},
				},
			},
		},
		bson.D{{"$set", bson.D{{"date", "$_id"}}}},
		bson.D{{"$project", bson.D{{"_id", 0}}}},
		bson.D{{"$out", "audit_log_sums"}},
	}

	_, err := mgm.Coll(&AuditLog{}).Aggregate(ctx, pipeline)
	return err

}
