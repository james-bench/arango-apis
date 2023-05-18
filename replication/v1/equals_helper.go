//
// DISCLAIMER
//
// Copyright 2022-2023 ArangoDB GmbH, Cologne, Germany
//

package v1

// Equals returns true when source and other have the same values.
func (source *DeploymentReplication_Status) Equals(other *DeploymentReplication_Status) bool {
	return source.GetPhase() == other.GetPhase() &&
		source.GetMessage() == other.GetMessage() &&
		source.GetShardsInSync() == other.GetShardsInSync() &&
		source.GetTotalShards() == other.GetShardsInSync() &&
		source.GetProgress() == other.GetProgress() &&
		source.GetSyncEndpoint() == other.GetSyncEndpoint() &&
		source.GetForwarderEndpoint() == other.GetForwarderEndpoint()
}
