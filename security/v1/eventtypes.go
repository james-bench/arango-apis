//
// DISCLAIMER
//
// Copyright 2019 ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// IPWhitelist event types

	// EventTypeIPWhitelistCreated is the type of event fired after an IP whitelist has been created
	// SubjectID contains the ID of the whitelist.
	EventTypeIPWhitelistCreated = "security.ipwhitelist̀.created"
	// EventTypeIPWhitelistUpdated is the type of event fired after an IP whitelist has been updated
	// SubjectID contains the ID of the whitelist.
	EventTypeIPWhitelistUpdated = "security.ipwhitelist̀.updated"
	// EventTypeIPWhitelistDeleted is the type of event fired after an IP whitelist has been (marked for) deleted
	// SubjectID contains the ID of the whitelist.
	EventTypeIPWhitelistDeleted = "security.ipwhitelist̀.deleted"
)
