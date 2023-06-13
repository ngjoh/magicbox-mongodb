package model

import (
	"github.com/kamva/mgm/v3"
)

type AuditLog struct {
	mgm.DefaultModel
	AppId   string `json:"appid"`
	Subject string `json:"subject"`
}

func LogAudit(app string, subject string) {

	newRecord := &AuditLog{

		AppId:   app,
		Subject: subject,
	}

	mgm.Coll(newRecord).Create(newRecord)

}
