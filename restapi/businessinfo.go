package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func getCountries() usecase.Interactor {
	type GetRequest struct {
	}

	type GetResponse struct {
		Countries []*model.Country `json:"countries"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := model.Countries()
		output.Countries = data

		return err

	})

	u.SetTitle("Get a country ")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		"Business Info",
	)
	return u
}

func getBusinessAndGroupUnits() usecase.Interactor {
	type GetRequest struct {
	}

	type GetResponse struct {
		Units []*model.BusinessGroupUnit `json:"units"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := model.BusinessGroupUnits()
		output.Units = data

		return err

	})

	u.SetTitle("Get busines and group units ")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		"Business Info",
	)
	return u
}
