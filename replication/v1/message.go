//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
//

package v1

import "encoding/json"

// Message holds any additional information that the user must know regarding
// the replication / migration process.
type Message struct {
	// Reason contains information related to the current state of the migration process
	// or any supporting information related to failures.
	Reason string `json:"reason,omitempty"`
	// DbErrors contains information related to database errors that could occur during the migration process.
	DbErrors string `json:"db_errors,omitempty"`
	// SyncErrors contains information related to errors from the ArangoSync component.
	SyncErrors string `json:"sync_errors,omitempty"`
}

// AsJSON returns a JSON representation of the `Message` struct.
func (m Message) AsJSON() string {
	// The error here will always be `nil` and hence we ignore it.
	// This is because we know `Message` will always contain string fields,
	// and will be handled correctly by the JSON encoder.
	// Otherwise, the caller may add redundant boilerplate code to handle an error that never occurs.
	b, _ := json.Marshal(m)
	return string(b)
}

// NewMessage returns a new Message type with the given reason, dbErrors and syncErrors strings.
func NewMessage(reason, dbErrors, syncErrors string) Message {
	return Message{
		Reason:     reason,
		DbErrors:   dbErrors,
		SyncErrors: syncErrors,
	}
}
