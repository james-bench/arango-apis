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
