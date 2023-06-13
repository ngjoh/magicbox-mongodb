package restapi

import (
	"context"
	"errors"

	"github.com/koksmat-com/koksmat/model"
	"github.com/swaggest/usecase"
	"github.com/swaggest/usecase/status"
)

var tag = "Shared Mailboxes"

func createSharedMailbox() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input model.SharedMailboxNewResponce, output *model.SharedMailboxNewResponce) error {

		result, err := model.CreateSharedMailbox(input.DisplayName, input.Alias, input.Name, input.Members, input.Owners, input.Readers)
		if err != nil {
			return err
		}

		*output = result
		return nil

	})

	u.SetTitle("Create a Shared Mailbox")
	u.SetDescription("Create a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func getSharedMailbox() usecase.Interactor {
	type GetRequest struct {
		Identity string `path:"id"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *[]model.SharedMailbox) error {

		return errors.New("Not implemented")

	})

	u.SetTitle("Get a shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func updateSharedMailbox() usecase.Interactor {
	type SharedMailboxUpdateRquest struct {
		Identity    string `path:"id"`
		DisplayName string `json:"displayName" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxUpdateRquest, output *model.SharedMailboxNewResponce) error {

		result, err := model.UpdateSharedMailbox(input.Identity, input.DisplayName)
		if err != nil {
			return err
		}

		*output = result
		return nil

	})
	u.SetTitle("Create a Shared Mailbox")
	u.SetDescription("Create a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func addSharedMailboxMembers() usecase.Interactor {
	type SharedMailboxAddMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"displayName" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailboxNewResponce) error {

		return errors.New("Not implemented")

	})
	u.SetTitle("Add members to a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func removeSharedMailboxMembers() usecase.Interactor {
	type SharedMailboxRemoveMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"displayName" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxRemoveMemberRquest, output *model.SharedMailboxNewResponce) error {

		return errors.New("Not implemented")

	})
	u.SetTitle("Removes members from a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func addSharedMailboxOwners() usecase.Interactor {
	type SharedMailboxAddMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"emails" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailboxNewResponce) error {

		return errors.New("Not implemented")

	})
	u.SetTitle("Add owners to a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func removeSharedMailboxOwners() usecase.Interactor {
	type SharedMailboxRemoveMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"emails" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxRemoveMemberRquest, output *model.SharedMailboxNewResponce) error {

		return errors.New("Not implemented")

	})
	u.SetTitle("Removes owners from a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func listSharedMailbox() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *[]model.SharedMailbox) error {
		result, err := model.GetSharedMailboxes()
		if err != nil {
			return err
		}

		*output = append(*output, result...)
		return nil

	})

	u.SetTitle("Get Shared Mailboxes")
	u.SetDescription("List all Shared Mailboxes")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func deleteSharedMailbox() usecase.Interactor {
	type DeleteRequest struct {
		Identity string `path:"id"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input DeleteRequest, output *[]model.SharedMailbox) error {

		return errors.New("Not implemented")

	})

	u.SetTitle("Delete a shared Mailboxes")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func addSharedMailboxEmail() usecase.Interactor {
	type SharedMailboxAddEmailRquest struct {
		Identity string `path:"id"`
		Email    string `json:"smtpaddress" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddEmailRquest, output *model.SharedMailboxNewResponce) error {

		return errors.New("Not implemented")

	})
	u.SetTitle("Add a smtp address to a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func removeSharedMailboxEmail() usecase.Interactor {
	type SharedMailboxRemoveEmailRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"emails" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxRemoveEmailRquest, output *struct{}) error {

		return errors.New("Not implemented")

	})
	u.SetTitle("Removes a smtp address from a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
