//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	"net/url"
	"path"

	rm "github.com/arangodb-managed/apis/resourcemanager/v1"
)

const (
	// KindInvoice is a constants for the kind of Invoice resources.
	KindInvoice = "Invoice"
)

// InvoiceURL creates a resource URL for the Invoice with given ID
// in given context.
func InvoiceURL(organizationID, invoiceID string) string {
	return path.Join(rm.OrganizationURL(organizationID), KindInvoice, url.PathEscape(invoiceID))
}
