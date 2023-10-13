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
	// Credit bundles permissions

	// PermissionCreditBundlesList is needed for listing credit bundles.
	PermissionCreditBundleList = "credit.creditbundle.list"
)

const (
	// Credit bundle usage permissions

	// PermissionCreditBundlesList is needed for listing credit bundle usages.
	PermissionCreditBundleUsageList = "credit.creditbundleusage.list"
)

const (
	// Credit usage report permissions

	// PermissionCreditUsageReportList is needed for listing credit usage reports.
	PermissionCreditUsageReportList = "credit.creditusagereport.list"
	// PermissionCreditUsageReportGet is needed for getting a single credit usage report.
	PermissionCreditUsageReportGet = "credit.creditusagereport.get"
)

const (
	// Credit usage projection permissions

	// PermissionCreditUsageProjectionGet is needed for getting credit usage projection.
	PermissionCreditUsageProjectionGet = "credit.creditbundleusageprojection.get"
)
