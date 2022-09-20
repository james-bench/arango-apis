//
// DISCLAIMER
//
// Copyright 2020 ArangoDB GmbH, Cologne, Germany
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
	// Notebook permissions

	// PermissionGetNotebook is required to get a notebook.
	PermissionGetNotebook = "notebook.notebook.get"
	// PermissionCreateNotebook is required to create a notebook.
	PermissionCreateNotebook = "notebook.notebook.create"
	// PermissionDeleteNotebook is required to delete a notebook.
	PermissionDeleteNotebook = "notebook.notebook.delete"
	// PermissionStopNotebook  is required to stop a notebook.
	PermissionShutdownNotebook = "notebook.notebook.shutdown"
	// PermissionPauseNotebook is required to pause a notebook.
	PermissionPauseNotebook = "notebook.notebook.pause"
	// PermissionResumeNotebook is required to pause a notebook.
	PermissionResumeNotebook = "notebook.notebook.resume"
	// PermissionListNotebooks is required to list notebooks.
	PermissionListNotebooks = "notebook.notebook.list"
)
