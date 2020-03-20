//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Author Robert Stam
//

package v1

const (
	// Network transfer destinations

	// NetworkTransferDestinationInternet specifies the destination (or source) from the ingress/egress as internet data transfer.
	NetworkTransferDestinationInternet = "Internet"
	// NetworkTransferDestinationSameRegion specifies the destination (or source) from the ingress/egress as inner region data transfer.
	NetworkTransferDestinationSameRegion = "SameRegion"
)
