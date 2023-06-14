package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func demo() usecase.Interactor {
	type DemoRequest struct {
		Hello string `json:"hello" binding:"required"`
	}
	type DemoResponse struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input DemoRequest, output *DemoResponse) error {

		return model.NewDemo(input.Hello)

	})

	u.SetTitle("Create Demo")

	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Demo")
	return u
}
