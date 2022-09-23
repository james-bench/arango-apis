//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
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

var (
	// Replication has started, waiting for sync masters / workers.
	DeploymentReplicationPhaseInitialising = "Initialising"
	// Sync masters / workers are ready, deployment is ready to start replication process.
	DeploymentReplicationPhaseInitialised = "Initialised"
	// Replication has started and currently in progress.
	DeploymentReplicationPhaseInProgress = "In-Progress"
	// All shards and collections are currently in-sync.
	DeploymentReplicationPhaseInSync = "In-Sync"
	// Replication is in an errored state.
	DeploymentReplicationPhaseError = "Error"
	// Replication could not complete successfully.
	DeploymentReplicationPhaseFailed = "Failed"
	// Replication is being stopped.
	DeploymentReplicationPhaseStopping = "Stopping"
	// Replication is stopped and all resources cleaned up properly.
	DeploymentReplicationPhaseCompleted = "Completed"
)
