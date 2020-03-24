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
// Author Ewout Prangsma
//

package auth

import (
	"fmt"
	"regexp"
	"strings"
)

var (
	partRegexp = regexp.MustCompile(`^[a-z][a-z\-]*$`)
)

// ParsePermission parses a permission into:
// - api
// - kind
// - verb
//
// E.g. `iam.group.get` returns `iam`, `group`, `get`.
func ParsePermission(p string) (string, string, string, error) {
	parts := strings.Split(p, ".")
	if len(parts) != 3 {
		return "", "", "", fmt.Errorf("Invalid permission syntax '%s'", p)
	}
	api, kind, verb := parts[0], parts[1], parts[2]
	if !partRegexp.MatchString(api) {
		return "", "", "", fmt.Errorf("Invalid permission (api) in '%s'", p)
	}
	if !partRegexp.MatchString(kind) {
		return "", "", "", fmt.Errorf("Invalid permission (kind) in '%s'", p)
	}
	if !partRegexp.MatchString(verb) {
		return "", "", "", fmt.Errorf("Invalid permission (verb) in '%s'", p)
	}
	return api, kind, verb, nil
}

// CreatePermission creates a permission from given arguments.
func CreatePermission(api, kind, verb string) string {
	return api + "." + kind + "." + verb
}
