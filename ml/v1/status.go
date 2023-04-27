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

import "sort"

const (
	// MLServices status phases

	// The services needed for ArangoGraphML are being installed.
	MLServicesPhaseInitialising = "Initialising"
	// ArangoDB Deployment is being bootstrapped with the required databases, schemas and data.
	MLServicesPhaseBootstrapping = "Bootstrapping"
	// ArangoGraphML is setup and running correctly.
	MLServicesPhaseRunning = "Running"
	// Indicates that there was an error with setting up ArangoGraphML. Check `message` field for additional info.
	MLServicesPhaseError = "Error"
)

const (
	// Service types

	// ServiceTypePrediction indicates that the service is a prediction API service.
	ServiceTypePrediction = "prediction"
	// ServiceTypeTraining indicates that the service is a training API service.
	ServiceTypeTraining = "training"
	// ServiceTypeProjects indicates that the service is a projects API service.
	ServiceTypeProjects = "projects"
)

// Equals returns true when source and other have the same values.
func (source *Status) Equals(other *Status) bool {
	return source.GetPhase() == other.GetPhase() &&
		source.GetMessage() == other.GetMessage() &&
		source.GetLastUpdatedAt().Equal(other.GetLastUpdatedAt()) &&
		Equals(source.GetServices(), other.GetServices())
}

// Equals returns true when source and other have the same values.
func (source *ServiceStatus) Equals(other *ServiceStatus) bool {
	return source.GetType() == other.GetType() &&
		source.GetAvailable() == other.GetAvailable() &&
		source.GetFailed() == other.GetFailed() &&
		source.GetReplicas() == other.GetReplicas() &&
		source.GetUsage().Equals(other.GetUsage())
}

// Equals returns true when source and other have the same values.
func Equals(source, other []*ServiceStatus) bool {
	if len(source) != len(other) {
		return false
	}

	// Sort both slices.
	sort.Slice(source, func(i, j int) bool {
		return source[i].GetType() < source[j].GetType()
	})
	sort.Slice(other, func(i, j int) bool {
		return other[i].GetType() < other[j].GetType()
	})

	// Check if every element is equal
	for i := range source {
		if !source[i].Equals(other[i]) {
			return false
		}
	}
	return true
}

// Equals returns true when source and other have the same values.
func (source *ServiceStatus_Usage) Equals(other *ServiceStatus_Usage) bool {
	return source.GetLastCpuUsage() == other.GetLastCpuUsage() &&
		source.GetLastCpuLimit() == other.GetLastCpuLimit() &&
		source.GetLastMemoryUsage() == other.GetLastMemoryUsage() &&
		source.GetLastMemoryLimit() == other.GetLastMemoryLimit()
}

// EnsureServiceStatus sets the status of a ML service.
func (status *Status) EnsureServiceStatus(svcStatus *ServiceStatus) {
	// Check if this service type already exists
	for i, s := range status.GetServices() {
		if s.GetType() == svcStatus.GetType() {
			status.Services[i] = svcStatus
			return
		}
	}
	// Otherwise, append to the services.
	status.Services = append(status.Services, svcStatus)
}
