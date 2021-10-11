//
// DISCLAIMER
//
// Copyright 2020-2021 ArangoDB GmbH, Cologne, Germany
//
// Author Robert Stam
//

package v1

const (
	// Network transfer destinations

	// NetworkTransferDestinationInternet specifies the destination (or source) from the ingress/egress as internet data transfer.
	NetworkTransferDestinationInternet = "Internet"
	// NetworkTransferDestinationInCluster specifies the destination (or source) from the ingress/egress as in-cluster data transfer.
	NetworkTransferDestinationInCluster = "InCluster"
	// NetworkTransferDestinationPrivateEndpoint specifies the destination (or source) from the ingress/egress as private-endpoint data transfer.
	NetworkTransferDestinationPrivateEndpoint = "PrivateEndpoint"
)
