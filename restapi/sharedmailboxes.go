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
	type SharedMailboxNewRequest struct {
		DisplayName string   `json:"displayName" binding:"required"`
		Alias       string   `json:"alias" binding:"required"`
		Name        string   `json:"name" binding:"required"`
		Members     []string `json:"members"`
		Owners      []string `json:"owners"`
		Readers     []string `json:"readers"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxNewRequest, output *model.SharedMailbox) error {

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
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *model.SharedMailbox) error {

		output, err := model.GetSharedMailbox(input.Identity)
		return err

	})

	u.SetTitle("Get a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func updateSharedMailbox() usecase.Interactor {
	type SharedMailboxUpdateRequest struct {
		Identity    string `path:"id"`
		DisplayName string `json:"displayName" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxUpdateRequest, output *model.SharedMailbox) error {
		result, err := model.UpdateSharedMailbox(input.Identity, input.DisplayName)
		*output = *result
		return err

	})
	u.SetTitle("Update a Shared Mailbox")
	u.SetDescription("Updates a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func updateSharedMailboxPrimaryEmailAddress() usecase.Interactor {

	type SharedMailboxUpdateRequest struct {
		Identity string `path:"id"`
		Email    string `json:"smtpaddress" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxUpdateRequest, output *model.SharedMailbox) error {

		result, err := model.UpdateSharedMailboxPrimaryEmailAddress(input.Identity, input.Email)
		*output = *result
		return err

	})
	u.SetTitle("Update a Shared Mailbox primary smtp address")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func addSharedMailboxMembers() usecase.Interactor {
	type SharedMailboxAddMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"members" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailbox) error {
		result, err := model.AddSharedMailboxMembers(input.Identity, input.Members)
		*output = *result
		return err

	})
	u.SetTitle("Add members to a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func removeSharedMailboxMembers() usecase.Interactor {
	type SharedMailboxRemoveMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"members" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxRemoveMemberRquest, output *model.SharedMailbox) error {

		result, err := model.RemoveSharedMailboxMembers(input.Identity, input.Members)
		*output = *result
		return err

	})
	u.SetTitle("Removes members from a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func addSharedMailboxReaders() usecase.Interactor {
	type SharedMailboxAddMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"readers" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailbox) error {
		result, err := model.AddSharedMailboxReaders(input.Identity, input.Members)
		*output = *result
		return err

	})
	u.SetTitle("Add readers to a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func removeSharedMailboxReaders() usecase.Interactor {
	type SharedMailboxRemoveMemberRquest struct {
		Identity string   `path:"id"`
		Members  []string `json:"readers" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxRemoveMemberRquest, output *model.SharedMailbox) error {

		output, err := model.RemoveSharedMailboxReaders(input.Identity, input.Members)
		return err

	})
	u.SetTitle("Removes readers from a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func setSharedMailboxOwners() usecase.Interactor {
	type SharedMailboxAddMemberRquest struct {
		Identity string   `path:"id"`
		Owners   []string `json:"owners" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailbox) error {
		result, err := model.SetSharedMailboxOwners(input.Identity, input.Owners)
		*output = *result
		return err

	})
	u.SetTitle("Add owners to a Shared Mailbox")
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
		return model.DeleteSharedMailbox(input.Identity)
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
	u.SetTitle("Add a smtp address to a Shared Mailbox [NOT IMPLEMENTED]")
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
	u.SetTitle("Removes a smtp address from a Shared Mailbox [NOT IMPLEMENTED]")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
