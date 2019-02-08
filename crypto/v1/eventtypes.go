//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// CACertificate event types

	// EventTypeCACertificateCreated is the type of event fired after a CA certificate has been created
	// SubjectID contains the ID of the CA certificate.
	EventTypeCACertificateCreated = "crypto.cacertificate.created"
	// EventTypeCACertificateDeleted is the type of event fired after a CA certificate has been (marked for) deleted
	// SubjectID contains the ID of the CA certificate.
	EventTypeCACertificateDeleted = "crypto.cacertificate.deleted"
)
