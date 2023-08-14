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
		DisplayName string   `json:"displayName" binding:"required" example:"Shared Mailbox Name"`
		Alias       string   `json:"alias" binding:"required" example:"shared-mailbox-alias" `
		Name        string   `json:"name" binding:"required" example:"shared-mailbox-name"`
		Members     []string `json:"members"`
		Owners      []string `json:"owners"`
		Readers     []string `json:"readers"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxNewRequest, output *model.SharedMailbox) error {

		result, err := model.CreateSharedMailbox(ctx.Value("auth").(model.Authorization), input.DisplayName, input.Alias, input.Name, input.Members, input.Owners, input.Readers)
		if err != nil {
			return err
		}

		*output = *result
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
		Identity string `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input GetRequest, output *model.SharedMailbox) error {

		data, err := model.GetSharedMailbox(ctx.Value("auth").(model.Authorization), input.Identity)
		*output = *data
		return err

	})

	u.SetTitle("Get a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func updateSharedMailbox() usecase.Interactor {
	type SharedMailboxUpdateRequest struct {
		Identity    string `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		DisplayName string `json:"displayName" binding:"required" example:"New Display Name for Shared Mailbox"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxUpdateRequest, output *model.SharedMailbox) error {
		result, err := model.UpdateSharedMailbox(ctx.Value("auth").(model.Authorization), input.Identity, input.DisplayName)
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
		Identity string `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		Email    string `json:"smtpaddress" binding:"required" example:"contact@contoso.com"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxUpdateRequest, output *model.SharedMailbox) error {

		result, err := model.UpdateSharedMailboxPrimaryEmailAddress(ctx.Value("auth").(model.Authorization), input.Identity, input.Email)
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
		Identity string   `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		Members  []string `json:"members" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailbox) error {
		result, err := model.AddSharedMailboxMembers(ctx.Value("auth").(model.Authorization), input.Identity, input.Members)
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
		Identity string   `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		Members  []string `json:"members" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxRemoveMemberRquest, output *model.SharedMailbox) error {

		result, err := model.RemoveSharedMailboxMembers(ctx.Value("auth").(model.Authorization), input.Identity, input.Members)
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
		Identity string   `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		Members  []string `json:"readers" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailbox) error {
		result, err := model.AddSharedMailboxReaders(ctx.Value("auth").(model.Authorization), input.Identity, input.Members)
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
		Identity string   `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		Members  []string `json:"readers" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxRemoveMemberRquest, output *model.SharedMailbox) error {

		data, err := model.RemoveSharedMailboxReaders(ctx.Value("auth").(model.Authorization), input.Identity, input.Members)
		*output = *data
		return err

	})
	u.SetTitle("Removes readers from a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}
func setSharedMailboxOwners() usecase.Interactor {
	type SharedMailboxAddMemberRquest struct {
		Identity string   `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		Owners   []string `json:"owners" binding:"required"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input SharedMailboxAddMemberRquest, output *model.SharedMailbox) error {
		result, err := model.SetSharedMailboxOwners(ctx.Value("auth").(model.Authorization), input.Identity, input.Owners)
		*output = *result
		return err

	})
	u.SetTitle("Set the owners of a Shared Mailbox")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func listSharedMailbox() usecase.Interactor {
	u := usecase.NewInteractor(func(ctx context.Context, input struct{}, output *[]model.SharedMailbox) error {

		result, err := model.GetSharedMailboxes(ctx.Value("auth").(model.Authorization))
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
		Identity string `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input DeleteRequest, output *[]model.SharedMailbox) error {
		return model.DeleteSharedMailbox(ctx.Value("auth").(model.Authorization), input.Identity)
	})

	u.SetTitle("Delete a shared Mailboxes")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags(tag)
	return u
}

func addSharedMailboxEmail() usecase.Interactor {
	type SharedMailboxAddEmailRquest struct {
		Identity string `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
		Email    string `json:"smtpaddress" binding:"required" example:"contact@contosoelectronics.com`
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
		Identity string   `path:"exchangeObjectId" example:"6ebef668-9d8b-4fe3-9d40-f641751f5944"`
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

func provisionRoom() usecase.Interactor {
	type RoomProvisionRequest struct {
		SharepointId int `json:"sharepointid" example:421`
	}
	type RoomProvisionResponse struct {
		Email string `json:"email" example:room-zzz@domain.com`
	}
	u := usecase.NewInteractor(func(ctx context.Context, input RoomProvisionRequest, output *RoomProvisionResponse) error {

		email, err := model.ProvisionRoomBySharePointID(input.SharepointId)
		output.Email = *email
		return err

	})
	u.SetTitle("Provision a room")
	u.SetDescription("Provision a room by referencing to a sharepoint item id")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Room Management (legacy)")
	return u
}

func deleteRoom() usecase.Interactor {
	type RoomDeleteRequest struct {
		SharepointId int `path:"sharepointitemid" example:421`
	}
	type RoomDeleteResponse struct {
	}
	u := usecase.NewInteractor(func(ctx context.Context, input RoomDeleteRequest, output *RoomDeleteResponse) error {

		_, err := model.DeleteRoomBySharePointID(input.SharepointId)

		return err

	})
	u.SetTitle("Deletes a room")
	u.SetDescription("Deletes a room by referencing to a sharepoint item id")
	u.SetExpectedErrors(status.InvalidArgument)
	u.SetTags("Room Management (legacy)")
	return u
}
