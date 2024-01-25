package restapi

import (
	"context"
	"encoding/json"

	"github.com/koksmat-com/koksmat/kitchen"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func GetMateJobs() usecase.Interactor {
	type MateRequest struct {
		ID         string   `json:"id" binding:"required" example:"1234567890"`
		Task       string   `json:"cmd" binding:"required" example:"kitchen"`
		Parameters []string `json:"parameters" binding:"required" example:["kitchen","status","365admin-test"] `
	}
	type MateResponse struct {
		ID           string `json:"id" binding:"required" example:"1234567890"`
		HasError     bool   `json:"hasError" binding:"required" example:"false"`
		ErrorMessage string `json:"errorMessage"  example:"Didn't work"`
		Response     string `json:"response" example:".."`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input MateRequest, output *MateResponse) error {

		switch input.Task {
		case "kitchen":
			switch input.Parameters[0] {
			case "list":
				k, err := kitchen.List()
				if err != nil {
					output.ErrorMessage = err.Error()
					output.HasError = true
					return nil
				}
				data, err := json.Marshal(k)
				if err != nil {
					if err != nil {
						output.ErrorMessage = err.Error()
						output.HasError = true
						return nil
					}
				}
				output.Response = string(data)
			default:
				output.ErrorMessage = "Unknown kitchen command"
				output.HasError = true
			}
		default:
			output.ErrorMessage = "Unknown task"
			output.HasError = true
		}
		return nil

	})

	u.SetTitle("Get Mate Jobs")
	u.SetDescription("")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Mate")
	return u
}
