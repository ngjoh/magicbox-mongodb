package io

import (
	"github.com/Azure/azure-sdk-for-go/sdk/storage/azblob"
)

func AzureSignin() (any, error) {
	return azblob.NewSharedKeyCredential("", "")

}
