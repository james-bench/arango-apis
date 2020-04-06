//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Author Gergely Brautigam
//

package v1

const (
	// Installation states

	// InProgress means that the pod was created successfull and the data is being generated / imported
	InProgress = "InProgress"
	// Created means that the data has been successfully generated.
	Created = "Created"
	// Ready means that the data is ready to be used.
	Ready = "Ready"
	// Failed denotes a dataset which has failed to be created.
	Failed = "Failed"
)
