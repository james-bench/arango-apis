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

package v1

import (
	fmt "fmt"
	"net/url"
	"path"
	"strings"
)

const (
	// KindOrganization is a constants for the kind of Organization resources.
	KindOrganization = "Organization"
	// KindProject is a constants for the kind of Project resources.
	KindProject = "Project"
	// KindOrganizationInvite is a constants for the kind of OrganizationInvite resources.
	KindOrganizationInvite = "OrganizationInvite"
)

// OrganizationURL creates a resource URL for the organization with given ID.
func OrganizationURL(organizationID string) string {
	return path.Join("/", KindOrganization, url.PathEscape(organizationID))
}

// ProjectURL creates a resource URL for the project & organization with given IDs.
func ProjectURL(organizationID, projectID string) string {
	return path.Join(OrganizationURL(organizationID), KindProject, url.PathEscape(projectID))
}

// OrganizationInviteURL creates a resource URL for the invite & organization with given IDs.
func OrganizationInviteURL(organizationID, inviteID string) string {
	return path.Join(OrganizationURL(organizationID), KindOrganizationInvite, url.PathEscape(inviteID))
}

// ResourceURL holds a parsed resource URL.
type ResourceURL []ResourceID

// String representation of the URL
func (u ResourceURL) String() string {
	builder := strings.Builder{}
	for _, x := range u {
		builder.WriteString("/")
		builder.WriteString(x.Kind)
		builder.WriteString("/")
		builder.WriteString(url.PathEscape(x.ID))
	}
	return builder.String()
}

// OrganizationID returns the organization ID of the given URL.
func (u ResourceURL) OrganizationID() string {
	return u[0].ID
}

// OrganizationChildID returns the ID of a child of the organization in the given URL.
// At least one kind must be given.
// If more kinds are given, they refer to a grant-child.
// If not (grant)child with matching kind(s) is found, an empty string is returned.
func (u ResourceURL) OrganizationChildID(kind ...string) string {
	var lastID string
	for i, k := range kind {
		idx := 1 + i
		if idx >= len(u) {
			return ""
		}
		if u[idx].Kind != k {
			return ""
		}
		lastID = u[idx].ID
	}
	return lastID
}

// ProjectID returns the project ID of the given URL.
// If the URL does not contain a project ID, an empty string is returned.
func (u ResourceURL) ProjectID() string {
	return u.OrganizationChildID(KindProject)
}

// ProjectChildID returns the ID of a child of the project in the given URL.
// At least one kind must be given.
// If more kinds are given, they refer to a grant-child.
// If not (grant)child with matching kind(s) is found, an empty string is returned.
func (u ResourceURL) ProjectChildID(kind ...string) string {
	var lastID string
	for i, k := range kind {
		idx := 2 + i
		if idx >= len(u) {
			return ""
		}
		if u[idx].Kind != k {
			return ""
		}
		lastID = u[idx].ID
	}
	return lastID
}

// Validate this URL.
// Returns nil is valid, error otherwise.
func (u ResourceURL) Validate() error {
	if len(u) == 0 {
		return fmt.Errorf("Empty resource URL")
	}
	if u[0].Kind != KindOrganization {
		return fmt.Errorf("Expected resource URL to start with '%s', got '%s'", KindOrganization, u[0].Kind)
	}
	for i, x := range u {
		if x.Kind == "" {
			return fmt.Errorf("Kind cannot be empty at index %d", i)
		}
		if x.ID == "" {
			return fmt.Errorf("ID cannot be empty at index %d", i)
		}
	}
	return nil
}

// ResourceID represents a single element of a resource URL hierarchy.
type ResourceID struct {
	Kind string
	ID   string
}

// ParseResourceURL parses the given resource URL into its elements.
func ParseResourceURL(resourceURL string) (ResourceURL, error) {
	if strings.HasPrefix(resourceURL, "/") {
		resourceURL = resourceURL[1:]
	}
	parts := strings.Split(resourceURL, "/")
	if len(parts)%2 != 0 {
		return nil, fmt.Errorf("Invalid resource URL: '%s'", resourceURL)
	}
	result := make(ResourceURL, 0, len(parts)/2)
	for i := 0; i < len(parts); i += 2 {
		kind := parts[i]
		id, err := url.PathUnescape(parts[i+1])
		if err != nil {
			return nil, fmt.Errorf("Invalid resource URL: '%s' (%s)", resourceURL, err)
		}
		result = append(result, ResourceID{
			Kind: kind,
			ID:   id,
		})
	}
	if err := result.Validate(); err != nil {
		return nil, err
	}
	return result, nil
}
