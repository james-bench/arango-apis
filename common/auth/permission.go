//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
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
