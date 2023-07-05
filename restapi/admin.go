package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/audit"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

const auditTag = "Audit"

func getAuditLogs() usecase.Interactor {
	type GetRequest struct {
		DateString string `path:"date" description:"date of the audit log" example:"2023-07-05"`
	}

	type GetResponse struct {
		AuditlogsSum    []*audit.AuditLogSum `json:"auditlogs"`
		NumberOfRecords int64                `json:"numberofrecords"`
		Pages           int64                `json:"pages"`
		CurrentPage     int64                `json:"currentpage"`
		PageSize        int64                `json:"pagesize"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := audit.GetAuditLogs(input.DateString)
		output.AuditlogsSum = data
		output.NumberOfRecords = int64(len(data))
		output.Pages = 1
		output.CurrentPage = 1
		output.PageSize = 100

		return err

	})

	u.SetTitle("Get audit logs ")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		auditTag,
	)
	return u
}

func getAuditLogPowershell() usecase.Interactor {
	type GetRequest struct {
		Id string `path:"objectId" description:"id of the audit log" example:"648dd75246669bf85c4d4e15"`
	}

	type GetResponse struct {
		PowerShellAuditlog *audit.PowerShellLog `json:"powershellauditlog"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := audit.GetPowerShellLog(input.Id)
		output.PowerShellAuditlog = data

		return err

	})

	u.SetTitle("Get audit logs ")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		auditTag,
	)
	return u
}
