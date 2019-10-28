//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	context "context"
)

type (
	// InvoiceCallback is a callback for individual invoices.
	InvoiceCallback func(context.Context, *Invoice) error

	// PaymentProviderCallback is a callback for individual payment providers.
	PaymentProviderCallback func(context.Context, *PaymentProvider) error

	// PaymentMethodCallback is a callback for individual payment methods.
	PaymentMethodCallback func(context.Context, *PaymentMethod) error
)

// ForEachInvoice iterates over all invoices that match the given request,
// invoking the given callback for each invoice.
func ForEachInvoice(ctx context.Context, c BillingServiceClient,
	req ListInvoicesRequest, cb InvoiceCallback) error {
	req.Options = req.Options.CloneOrDefault()
	for {
		list, err := c.ListInvoices(ctx, &req)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(req.Options.PageSize) {
			// We're done
			return nil
		}
		req.Options.Page++
	}
}

// ForEachPaymentProvider iterates over all payment providers that match the given request,
// invoking the given callback for each provider.
func ForEachPaymentProvider(ctx context.Context, c BillingServiceClient,
	req ListPaymentProvidersRequest, cb PaymentProviderCallback) error {
	req.Options = req.Options.CloneOrDefault()
	for {
		list, err := c.ListPaymentProviders(ctx, &req)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(req.Options.PageSize) {
			// We're done
			return nil
		}
		req.Options.Page++
	}
}

// ForEachPaymentMethod iterates over all payment methods that match the given request,
// invoking the given callback for each method.
func ForEachPaymentMethod(ctx context.Context, c BillingServiceClient,
	req ListPaymentMethodsRequest, cb PaymentMethodCallback) error {
	req.Options = req.Options.CloneOrDefault()
	for {
		list, err := c.ListPaymentMethods(ctx, &req)
		if err != nil {
			return err
		}
		for _, item := range list.GetItems() {
			if err := cb(ctx, item); err != nil {
				return err
			}
			if err := ctx.Err(); err != nil {
				return err
			}
		}
		if len(list.GetItems()) < int(req.Options.PageSize) {
			// We're done
			return nil
		}
		req.Options.Page++
	}
}