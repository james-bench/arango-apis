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
	// Token event types

	// EventTypeTokenCreated is the type of event fired after a metrics token has been created
	// SubjectID contains the ID of the token.
	EventTypeTokenCreated = "metrics.token.created"
	// EventTypeTokenUpdated is the type of event fired after a metrics token has been updated
	// SubjectID contains the ID of the token.
	EventTypeTokenUpdated = "metrics.token.updated"
	// EventTypeTokenDeleted is the type of event fired after a metrics token has been (marked for) deleted
	// SubjectID contains the ID of the token.
	EventTypeTokenDeleted = "metrics.token.deleted"
)
