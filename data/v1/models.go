//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// Deployment models

	// ModelOneShard is the model that yields a deployment with:
	// - a fixed number of nodes (3).
	// - a variable node size.
	// - a limitation of 1 shard per collection.
	// - a minimum replication factor of 2 (per collection)
	// - optimized for graph use cases
	ModelOneShard = "oneshard"

	// ModelSharded is the model that yields a deployment with:
	// - a variable number of nodes (3..).
	// - a variable node size.
	// - a minimum replication factor of 2 (per shard)
	// - no limitation of shards per collection
	ModelSharded = "sharded"

	// ModelFlexible is the model that yields a deployment with:
	// - completely flexible in number of coordinators, dbservers
	// - completely flexible in resources used by coordinators, dbservers
	// - no limitation of shards per collection
	// - a minimum replication factor of 1 (per collection)
	ModelFlexible = "flexible"
)
