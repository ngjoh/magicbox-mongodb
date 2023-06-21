package restapi

import (
	"context"
	"net/url"

	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

var tagAddresses = "Addresses"

func resolveAddress() usecase.Interactor {
	type GetRequest struct {
		Address string `path:"address"`
		Max     int    `query:"max" default:"10"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *model.Recipient) error {
		address, err := url.QueryUnescape(input.Address)
		if err != nil {
			return err
		}
		result, err := model.FindRecipientByAddress(address)
		*output = *result
		return err
	})

	u.SetTitle("Lookup an address")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tagAddresses)
	return u
}
