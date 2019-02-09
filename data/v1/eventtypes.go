//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
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
