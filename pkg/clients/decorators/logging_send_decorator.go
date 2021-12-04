// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT license.

package decorators

import (
	"net/http"

	"github.com/Azure/go-autorest/autorest"

	"github.com/Azure/aks-deployer/pkg/clients/httpclient"
)

// DoLogging get the logging decorator
func DoLogging(region string) autorest.SendDecorator {
	return func(s autorest.Sender) autorest.Sender {
		return autorest.SenderFunc(
			httpclient.InstrumentWithConnection(
				&roundTripperShim{Sender: s},
				region,
			).RoundTrip,
		)
	}
}

// roundTripperShim converts an autorest.Sender to http.RoundTripper.
type roundTripperShim struct {
	Sender autorest.Sender
}

func (r *roundTripperShim) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.Sender.Do(req)
}
