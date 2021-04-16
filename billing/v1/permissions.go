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

const (
	// BillingConfig permissions

	// PermissionConfigGet is needed for getting the billing configuration of an organization
	PermissionConfigGet = "billing.config.get"
	// PermissionConfigSet is needed for setting the billing configuration of an organization
	PermissionConfigSet = "billing.config.set"
)

const (
	// Organization permissions

	// PermissionOrganizationGet is needed for getting the billing flags of an organization
	PermissionOrganizationGet = "billing.organization.get"
)

const (
	// Invoice permissions

	// PermissionInvoiceList is needed for listing invoices in an organization
	PermissionInvoiceList = "billing.invoice.list"
	// PermissionInvoiceGet is needed for getting indivual invoices in an organization
	PermissionInvoiceGet = "billing.invoice.get"
	// PermissionInvoiceGetPremiminary is needed for getting the preliminary invoice in an organization
	PermissionInvoiceGetPremiminary = "billing.invoice.get-preliminary"
)

const (
	// PaymentProvider permissions

	// PermissionPaymentProviderList is needed for listing payment providers in the context of an organization
	PermissionPaymentProviderList = "billing.paymentprovider.list"
)

const (
	// PaymentMethod permissions

	// PermissionPaymentMethodList is needed for listing payment methods created for an organization
	PermissionPaymentMethodList = "billing.paymentmethod.list"
	// PermissionPaymentMethodGet is needed for getting individual payment methods created for an organization
	PermissionPaymentMethodGet = "billing.paymentmethod.get"
	// PermissionPaymentMethodDelete is needed for deleting individual payment methods
	PermissionPaymentMethodDelete = "billing.paymentmethod.delete"
	// PermissionPaymentMethodCreate is needed for creating individual payment methods
	PermissionPaymentMethodCreate = "billing.paymentmethod.create"
	// PermissionPaymentMethodUpdate is needed for updating individual payment methods
	PermissionPaymentMethodUpdate = "billing.paymentmethod.update"
	// PermissionPaymentMethodGetDefault is needed for getting the default payment method for an organization
	PermissionPaymentMethodGetDefault = "billing.paymentmethod.get-default"
	// PermissionPaymentMethodSetDefault is needed for changing the default payment method for an organization
	PermissionPaymentMethodSetDefault = "billing.paymentmethod.set-default"
)
