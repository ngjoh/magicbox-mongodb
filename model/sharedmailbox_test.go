package model

import (
	//"log"
	"fmt"
	"strings"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func auth() Authorization {
	return Authorization{AppId: "test", Permissions: "*"}
}

func TestSharedMailbox(t *testing.T) {

	t.Log("Creating Shared Mailbox")
	id := fmt.Sprintf("test-%s", uuid.New())
	mb, err := CreateSharedMailbox(auth(), id, id, id, []string{"test"}, []string{"test"}, []string{"test"})
	if err != nil {
		t.Fatalf("Failed to create shared mailbox: %s", err)
	}

	assert.NotNil(t, mb.ExchangeObjectId)
	t.Log("Reading Shared Mailbox")
	mb2, err := GetSharedMailbox(auth(), mb.ExchangeObjectId)
	if err != nil {
		t.Fatalf("Failed to get shared mailbox: %s", err)
	}
	t.Log("Changing name Shared Mailbox")
	assert.Equal(t, mb2.DisplayName, id)
	newName := "New Name"
	_, err = UpdateSharedMailbox(auth(), mb.ExchangeObjectId, newName)
	if err != nil {
		t.Fatalf("Failed to update shared mailbox: %s", err)
	}
	t.Log("Validating name change")
	mb2, err = GetSharedMailbox(auth(), mb.ExchangeObjectId)
	if err != nil {
		t.Fatalf("Failed to get shared mailbox: %s", err)
	}
	assert.Equal(t, mb2.DisplayName, newName)

	t.Log("Deleting Shared Mailbox")
	err = DeleteSharedMailbox(auth(), mb.ExchangeObjectId)
	if err != nil {
		t.Fatalf("Failed to delete shared mailbox: %s", err)
	}
	t.Log("Done")
}

func TestMemberships(t *testing.T) {

	t.Log("Creating Shared Mailbox")
	id := fmt.Sprintf("test-%s", uuid.New())
	mb, err := CreateSharedMailbox(auth(), id, id, id, []string{}, []string{}, []string{})
	if err != nil {
		t.Fatalf("Failed to create shared mailbox: %s", err)
	}

	assert.NotNil(t, mb.ExchangeObjectId)

	t.Log("Adding members")
	mb2, err := AddSharedMailboxMembers(auth(), mb.ExchangeObjectId, []string{"alexw"})
	if err != nil {
		t.Fatalf("Failed to add members: %s", err)
	}

	assert.True(t, strings.HasPrefix(mb2.Members[0], "AlexW"), "Should have AlexW as member")

	t.Log("Adding additional members")
	mb2, err = AddSharedMailboxMembers(auth(), mb.ExchangeObjectId, []string{"ChristieC"})
	if err != nil {
		t.Fatalf("Failed to add members: %s", err)
	}

	assert.Equal(t, len(mb2.Members), 2, "Should be two members")

	t.Log("Removing a members")
	mb2, err = RemoveSharedMailboxMembers(auth(), mb.ExchangeObjectId, []string{"alexw"})
	if err != nil {
		t.Fatalf("Failed to remove members: %s", err)
	}

	assert.Equal(t, len(mb2.Members), 1, "Should be one member")

	t.Log("Adding readers")
	mb2, err = AddSharedMailboxReaders(auth(), mb.ExchangeObjectId, []string{"alexw"})
	if err != nil {
		t.Fatalf("Failed to add readers: %s", err)
	}

	assert.True(t, strings.HasPrefix(mb2.Members[0], "AlexW"), "Should have AlexW as member")

	t.Log("Adding additional readers")
	mb2, err = AddSharedMailboxReaders(auth(), mb.ExchangeObjectId, []string{"ChristieC"})
	if err != nil {
		t.Fatalf("Failed to add readers: %s", err)
	}

	assert.Equal(t, len(mb2.Members), 2, "Should be two readers")

	t.Log("Removing a reader")
	mb2, err = RemoveSharedMailboxReaders(auth(), mb.ExchangeObjectId, []string{"alexw"})
	if err != nil {
		t.Fatalf("Failed to remove readers: %s", err)
	}

	assert.Equal(t, len(mb2.Members), 1, "Should be one reader")

	t.Log("Deleting Shared Mailbox")
	err = DeleteSharedMailbox(auth(), mb.ExchangeObjectId)
	if err != nil {
		t.Fatalf("Failed to delete shared mailbox: %s", err)
	}
	t.Log("Done")
}
