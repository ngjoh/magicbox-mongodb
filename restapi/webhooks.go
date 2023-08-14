package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/officegraph"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

const webhooksTag = "Webhooks"

func getWebHooks() usecase.Interactor {
	type GetRequest struct {
		//	Paging `bson:",inline"`
	}

	type GetResponse struct {
		Webhooks []*officegraph.MicrosoftGraphSubscription `json:"webhooks"`
		// NumberOfRecords int64                                     `json:"numberofrecords"`
		// Pages           int64                                     `json:"pages"`
		// CurrentPage     int64                                     `json:"currentpage"`
		// PageSize        int64                                     `json:"pagesize"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := officegraph.SubscriptionList()
		output.Webhooks = data
		// output.NumberOfRecords = int64(len(data))
		// output.Pages = 1
		// output.CurrentPage = 1
		// output.PageSize = 100

		return err

	})

	u.SetTitle("Get webhooks ")

	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		webhooksTag,
	)
	return u
}
