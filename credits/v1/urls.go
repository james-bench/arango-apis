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

import (
	"net/url"
	"path"
)

const (
	// KindCreditBundle is a constant for the kind of CreditBundle resources.
	KindCreditBundle = "CreditBundle"
	// KindCreditUsageReport is a constant for the kind of CreditUsageReport resources.
	KindCreditUsageReport = "CreditUsageReport"
)

// CreditBundleURL creates a resource URL for a credit bundle with the given ID.
func CreditBundleURL(organizationURL, bundleID string) string {
	return path.Join(organizationURL, KindCreditBundle, url.PathEscape(bundleID))
}

func CreditUsageReportURL(organizationURL, reportID string) string {
	return path.Join(organizationURL, KindCreditUsageReport, url.PathEscape(reportID))
}
