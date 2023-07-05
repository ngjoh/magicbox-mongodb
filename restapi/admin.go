package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/audit"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

const auditTag = "Audit"

type Paging struct {
	Page     int64 `query:"page" description:"page number" example:"1"`
	PageSize int64 `query:"pagesize" description:"page size" example:"100"`
}

func GetAuditLogSummarys() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *[]audit.AuditLogSum) error {
		result, err := audit.GetAuditLogSummarys()
		if err != nil {
			return err
		}
		for _, item := range result {

			*output = append(*output, *item)
		}
		return nil

	})

	u.SetTitle("Get Audit Log summary")
	u.SetDescription("Get Audit Log summary")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(auditTag)
	return u
}
func getAuditLogs() usecase.Interactor {
	type GetRequest struct {
		Paging     `bson:",inline"`
		DateString string `path:"date" description:"date of the audit log" example:"2023-07-05"`
		HourString string `path:"hour" description:"hour of the audit log" example:"09"`
	}

	type GetResponse struct {
		AuditlogsSum    []*audit.PowerShellLog `json:"auditlogs"`
		NumberOfRecords int64                  `json:"numberofrecords"`
		Pages           int64                  `json:"pages"`
		CurrentPage     int64                  `json:"currentpage"`
		PageSize        int64                  `json:"pagesize"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := audit.GetAuditLogs(input.DateString, input.HourString)
		output.AuditlogsSum = data
		output.NumberOfRecords = int64(len(data))
		output.Pages = 1
		output.CurrentPage = 1
		output.PageSize = 100

		return err

	})

	u.SetTitle("Get audit logs ")
	u.SetDescription("Get audit logs by date and hour - timezone is  GMT")
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
