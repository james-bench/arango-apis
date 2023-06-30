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
	// Deployment event types

	// EventTypeDeploymentCreated is the type of event fired after a deployment has been created
	// SubjectID contains the ID of the deployment.
	EventTypeDeploymentCreated = "data.deployment.created"
	// EventTypeDeploymentUpdated is the type of event fired after a deployment has been updated
	// SubjectID contains the ID of the deployment.
	// Note that this type of event is also fired when the status of the deployment has changed.
	EventTypeDeploymentUpdated = "data.deployment.updated"
	// EventTypeDeploymentDeleted is the type of event fired after a deployment has been (marked for) deleted
	// SubjectID contains the ID of the deployment.
	EventTypeDeploymentDeleted = "data.deployment.deleted"
)

const (
	// DeploymentCredentials event types

	// EventTypeDeploymentCredentialsRead is the type of event fired after the credentials
	// of a deployment have been read.
	// SubjectID contains the ID of the deployment.
	EventTypeDeploymentCredentialsRead = "data.deploymentcredentials.read"
)

const (
	// Test database event types

	// EventTypeDeploymentCreatedTestDatabase is the type of event fired after a test database has been created
	// SubjectID contains the ID of the deployment.
	// Contains the following custom fields:
	// username - the username of the created user
	// database - the name of the created database
	EventTypeDeploymentCreatedTestDatabase = "data.deployment.created-test-database"
)

const (
	// Deployment event reasons

	// DeploymentEventReasonPauseDeployment contains the reason that should be set
	// inside the event that is emitted when a deployment is paused.
	DeploymentEventReasonPauseDeployment = "pause_deployment"
	// DeploymentEventReasonResumeDeployment contains the reason that should be set
	// inside the event that is emitted when a deployment is resumed.
	DeploymentEventReasonResumeDeployment = "resume_deployment"
)
