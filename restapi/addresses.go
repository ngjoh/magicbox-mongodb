package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

var tagAddresses = "Addresses"

func resolveAddress() usecase.Interactor {
	type GetRequest struct {
		Address string `path:"address"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *model.Recipient) error {

		output, err := model.FindRecipientByAddress(input.Address)

		return err
	})

	u.SetTitle("Lookup an address")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tagAddresses)
	return u
}
