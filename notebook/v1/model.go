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

// Model defines the configuration for a notebook.
type Model struct {
	// CPU units required by the notebook.
	// 1 CPU unit equals 1 physical / virtual CPU.
	CPU int32
	// Memory required by the notebook in 'GiB'.
	Memory int32
}

// GetCPU returns the CPU specific of the model.
func (m Model) GetCPU() int32 {
	return m.CPU
}

// GetMemory returns the Memory specification of the model.
func (m Model) GetMemory() int32 {
	return m.Memory
}

// Models is a map of all supported models on Oasis.
var Models = map[string]Model{
	"basic": {
		CPU:    2,
		Memory: 4,
	},
}
