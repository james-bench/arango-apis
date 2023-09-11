//
// DISCLAIMER
//
// Copyright 2023 ArangoDB GmbH, Cologne, Germany
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

package v1

const (
	// DeploymentLowerPort is the lower port number for the deployment, used with well-known certificates
	DeploymentLowerPort = 8529
	// DeploymentHigherPort is the higher port number for the deployment, used with self-signed certificates
	DeploymentHigherPort = 18529
	// ExternalServicesLowerPort is the lower port number for the external services, used with well-known certificates
	ExternalServicesLowerPort = 8829
	// ExternalServicesHigherPort is the higher port number for the external services, used with self-signed certificates
	ExternalServicesHigherPort = 18829
)
