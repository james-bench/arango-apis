//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
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
	organizationKind = "Organization"
	projectKind      = "Project"
)

// OrganizationURL creates a resource URL for the organization with given ID.
func OrganizationURL(organizationID string) string {
	return path.Join("/", organizationKind, url.PathEscape(organizationID))
}

// ProjectURL creates a resource URL for the project & organization with given IDs.
func ProjectURL(organizationID, projectID string) string {
	return path.Join(OrganizationURL(organizationID), projectKind, url.PathEscape(projectID))
}

// ResourceURL holds a parsed resource URL.
type ResourceURL []ResourceID

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
	return u.OrganizationChildID(projectKind)
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
	if u[0].Kind != organizationKind {
		return fmt.Errorf("Expected resource URL to start with '%s', got '%s'", organizationKind, u[0].Kind)
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
