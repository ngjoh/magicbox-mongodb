// -------------------------------------------------------------------
// Generated by 365admin-publish/api/20 makeschema.ps1
// -------------------------------------------------------------------
/*
---
title: Backup all databases
---
*/
package endpoints

import (
	"context"
	"encoding/json"
	"os"
	"path"

	"github.com/swaggest/usecase"

	"github.com/365admin/magicbox-mongodb/execution"
	"github.com/365admin/magicbox-mongodb/schemas"
	"github.com/365admin/magicbox-mongodb/utils"
)

func BackupAllPost() usecase.Interactor {
	type Request struct {
		Body schemas.Databaseservices `json:"body" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input Request, output *string) error {
		body, inputErr := json.Marshal(input.Body)
		if inputErr != nil {
			return inputErr
		}

		inputErr = os.WriteFile(path.Join(utils.WorkDir("magicbox-mongodb"), "databaseservices.json"), body, 0644)
		if inputErr != nil {
			return inputErr
		}

		_, err := execution.ExecutePowerShell("john", "*", "magicbox-mongodb", "10-backup", "20 backup-all.ps1", "")
		if err != nil {
			return err
		}

		return err

	})
	u.SetTitle("Backup all databases")
	// u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Backup")
	return u
}
