//
// DISCLAIMER
//
// Copyright 2020-2023 ArangoDB GmbH, Cologne, Germany
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

package v1

const (
	// Formatting options

	// FormatText defines the log format to be text
	FormatText = "text"
	// FormatJson defines the log format to be json.
	// In JSON format, log lines are encoded as a single line JSON object.
	FormatJSON = "json"
)

const (
	// MetricsServerTypeDBServer should be set in GetDeploymentUsageMetricsRequest to indicate
	// that metrics are being requested for DBServers
	MetricsServerTypeDBServer = "dbserver"
	// MetricsServerTypeCoordinator should be set in GetDeploymentUsageMetricsRequest to indicate
	// that metrics are being requested for Coordinators
	MetricsServerTypeCoordinator = "coordinator"
	// MetricsServerTypeSingle should be set in GetDeploymentUsageMetricsRequest to indicate
	// that metrics are being requested for Single server
	MetricsServerTypeSingle = "single"

	// MetricTypeCPU should be set in GetDeploymentUsageMetricsRequest when CPU metrics
	// are being requested
	MetricTypeCPU = "cpu"
	// MetricTypeMemory should be set in GetDeploymentUsageMetricsRequest when Memory metrics
	// are being requested
	MetricTypeMemory = "memory"
	// MetricTypeDisk should be set in GetDeploymentUsageMetricsRequest when Disk metrics
	// are being requested
	MetricTypeDisk = "disk"
)
