//
// DISCLAIMER
//
// Copyright 2021 ArangoDB GmbH, Cologne, Germany
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
// Author Robert Stam
//

package v1

const (
	// PrivateEndpointService event types

	// EventTypePrivateEndpointServiceCreated is the type of event fired after a private endpoint service has been created
	// SubjectID contains the ID of the private endpoint service.
	EventTypePrivateEndpointServiceCreated = "network.privateendpointservice.created"
	// EventTypePrivateEndpointServiceUpdated is the type of event fired after a private endpoint service has been updated
	// SubjectID contains the ID of the private endpoint service.
	// Note that this type of event is also fired when the status of the private endpoint service has changed.
	EventTypePrivateEndpointServiceUpdated = "network.privateendpointservice.updated"
	// EventTypePrivateEndpointServiceDeleted is the type of event fired after a private endpoint service has been (marked for) deleted
	// SubjectID contains the ID of the private endpoint service.
	EventTypPrivateEndpointServiceDeleted = "network.privateendpointservice.deleted"
)
