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
	// Notebook related quotas

	// QuotaKindNotebooksPerDeployment limits the number of notebooks in a single (non-free) deployment.
	// This kind of quota must be requested on a project level.
	QuotaKindNotebooksPerDeployment = "notebook.total-notebook-per-deployment"

	// QuotaKindNotebooksPerFreeDeployment limits the number of notebooks in a single free deployment.
	// This kind of quota must be requested on a project level.
	QuotaKindNotebooksPerFreeDeployment = "notebook.total-notebook-per-free-deployment"
)
