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

const (
	// ML quotas

	// QuotaKindMLTrialJobUsageHours limits the total number of hours ML jobs can be run for,
	// for free deployments.
	// This kind of quota must be requested on a project level.
	QuotaKindMLTrialJobUsageHours = "ml.trial-usage-hours"

	// QuotaKindMLTrialExpiryDays limits the number of days ML may be enabled for free deployments.
	// This kind of quota must be requested on a project level.
	QuotaKindMLTrialExpiryDays = "ml.trial-expiry-days"
)
