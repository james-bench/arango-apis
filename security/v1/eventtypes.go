//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// Copyright holder is ArangoDB GmbH, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// IPAllowlist event types

	// EventTypeIPAllowlistCreated is the type of event fired after an IP allowlist has been created
	// SubjectID contains the ID of the allowlist.
	EventTypeIPAllowlistCreated = "security.ipallowlist.created"
	// EventTypeIPAllowlistUpdated is the type of event fired after an IP allowlist has been updated
	// SubjectID contains the ID of the allowlist.
	EventTypeIPAllowlistUpdated = "security.ipallowlist.updated"
	// EventTypeIPAllowlistDeleted is the type of event fired after an IP allowlist has been (marked for) deleted
	// SubjectID contains the ID of the allowlist.
	EventTypeIPAllowlistDeleted = "security.ipallowlist.deleted"
)

const (
	// IPWhitelist event types
	// Note: The use of these constants has been deprecated.
	// In a future version, they will be removed.

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

const (
	// IAMProvider event types

	// EventTypeIAMProviderCreated is the type of event fired after an IAM provider has been created
	// SubjectID contains the ID of the provider.
	EventTypeIAMProviderCreated = "security.iamprovider.created"
	// EventTypeIAMProviderUpdated is the type of event fired after an IAM provider has been updated
	// SubjectID contains the ID of the provider.
	EventTypeIAMProviderUpdated = "security.iamprovider.updated"
	// EventTypeIAMProviderDeleted is the type of event fired after an IAM provider has been (marked for) deleted
	// SubjectID contains the ID of the provider.
	EventTypeIAMProviderDeleted = "security.iamprovider.deleted"
)
