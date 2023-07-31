package restapi

import (
	"context"

	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

func getUsers() usecase.Interactor {
	type GetRequest struct {
	}

	type GetResponse struct {
		Users []*model.User `json:"users"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *GetResponse) error {

		data, err := model.GetUsers(ctx.Value("auth").(model.Authorization))
		output.Users = data

		return err

	})

	u.SetTitle("Get all users")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		"Users",
	)
	return u
}

func addUser() usecase.Interactor {
	type AddRequest struct {
		UPN         string `json:"upn" binding:"required" example:"jd@domain.com"	`
		DisplayName string `json:"displayName" binding:"required" example:"John Doe"	`
	}

	type AddResponse struct {
		User *model.User `json:"user"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input AddRequest, output *AddResponse) error {

		data, err := model.CreateUser(ctx.Value("auth").(model.Authorization), input.UPN, input.DisplayName)
		output.User = data

		return err

	})

	u.SetTitle("Create an user ")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		"Users",
	)
	return u
}

func updateUserCredentials() usecase.Interactor {

	type UpdateRequest struct {
		UPN         string             `path:"upn"  example:"jd@domain.com"	`
		Credentials []model.Credential `json:"credentials" binding:"required"`
	}

	type UpdateResponse struct {
		User *model.User `json:"user"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input UpdateRequest, output *UpdateResponse) error {

		data, err := model.UpdateUserCredentials(ctx.Value("auth").(model.Authorization), input.UPN, input.Credentials)
		output.User = data

		return err

	})

	u.SetTitle("Update the credentials of an user ")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(
		"Users",
	)
	return u
}
