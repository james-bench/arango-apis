//
// DISCLAIMER
//
// Copyright 2020-2021 ArangoDB GmbH, Cologne, Germany
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
	// Token permissions

	// PermissionTokenList is needed for listing metrics tokens in a deployment
	PermissionTokenList = "metrics.token.list"
	// PermissionTokenGet is needed for fetching an individual metrics token in a deployment
	PermissionTokenGet = "metrics.token.get"
	// PermissionTokenCreate is needed for create a metrics token
	PermissionTokenCreate = "metrics.token.create"
	// PermissionTokenUpdate is needed for updating a metrics token
	PermissionTokenUpdate = "metrics.token.update"
	// PermissionTokenRevoke is needed for revoking a metrics token
	PermissionTokenRevoke = "metrics.token.revoke"
	// PermissionTokenDelete is needed for deleting a metrics token
	PermissionTokenDelete = "metrics.token.delete"
)

const (
	// Metrics endpoint permissions

	// PermissionEndpointGet is needed for fetching the metrics endpoint for a deployment
	PermissionEndpointGet = "metrics.endpoint.get"
)
