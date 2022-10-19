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
	// Reason for any failures / conditions.
	Reason string `json:"reason,omitempty"`
	// Errors related to database that may occur during replication.
	DbErrors string `json:"db_errors,omitempty"`
	// Errors related to ArangoSync that may occur during replication.
	SyncErrors string `json:"sync_errors,omitempty"`
}

// AsJSON returns a JSON representation of the `Message` struct.
func (m Message) AsJSON() string {
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
