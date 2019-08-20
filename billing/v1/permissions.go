//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// Invoice permissions

	// PermissionInvoiceList is needed for listing invoices in an organization
	PermissionInvoiceList = "billing.invoice.list"
	// PermissionInvoiceGet is needed for getting indivual invoices in an organization
	PermissionInvoiceGet = "billing.invoice.get"
)
