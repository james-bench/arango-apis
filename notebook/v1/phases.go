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
	// Notebook status phases

	// Notebook is initialising.
	NotebookPhaseInitialising = "Initialising"

	// Notebook is running.
	NotebookPhaseRunning = "Running"

	// Notebook is moving to a hibernated state.
	NotebookPhaseHibernating = "Hibernating"

	// Notebook has moved to a hibernated state.
	NotebookPhaseHibernated = "Hibernated"

	// Notebook is in an errored state. Additional information can be obtained from `message` field.
	NotebookPhaseError = "Error"

	// Notebook has been marked for deletion and will clean-up all related resources.
	NotebookPhaseDeleting = "Deleting"
)
