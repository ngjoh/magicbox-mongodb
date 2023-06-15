package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/powershell"
	"github.com/spf13/viper"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func getInfo() usecase.Interactor {
	type InfoRequest struct {
	}
	type InfoResponse struct {
		Version string `json:"version" binding:"required"`
		Tenant  string `json:"tenant" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input InfoRequest, output *InfoResponse) error {

		*&output.Version = "0.0.1"
		*&output.Tenant = viper.GetString("DATABASE")
		return nil

	})

	u.SetTitle("Get runtime info")

	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Info")
	return u
}

func getDomains() usecase.Interactor {
	type InfoRequest struct {
	}

	u := usecase.NewInteractor(func(ctx context.Context, input InfoRequest, output *[]powershell.Domain) error {
		domains, err := powershell.GetDomains()
		*output = *domains
		return err

	})

	u.SetTitle("Get supported domains")

	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Info")
	return u
}
