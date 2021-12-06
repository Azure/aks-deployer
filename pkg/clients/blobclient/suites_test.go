package blobclient

import (
	"github.com/Azure/go-autorest/autorest"
)

func newMockAutoRestClient(sender autorest.Sender) autorest.Client {
	return autorest.Client{
		Sender: sender,
	}
}
