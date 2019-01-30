//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// CACertificate permissions

	// PermissionCACertificateList is needed for listing CA certificates in a project
	PermissionCACertificateList = "crypto.cacertificate.list"
	// PermissionCACertificateGet is needed for fetching an individual CA certificates in a project
	PermissionCACertificateGet = "crypto.cacertificate.get"
	// PermissionCACertificateCreate is needed for create a CA certificate
	PermissionCACertificateCreate = "crypto.cacertificate.create"
	// PermissionCACertificateUpdate is needed for updating a CA certificate
	PermissionCACertificateUpdate = "crypto.cacertificate.update"
	// PermissionCACertificateDelete is needed for deleting a CA certificate
	PermissionCACertificateDelete = "crypto.cacertificate.delete"
)
