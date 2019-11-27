//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'
import { IDOptions as arangodb_cloud_common_v1_IDOptions } from '../../common/v1/common'
import { ListOptions as arangodb_cloud_common_v1_ListOptions } from '../../common/v1/common'

// File: billing/v1/billing.proto
// Package: arangodb.cloud.billing.v1

// Address of organization
export interface Address {
  // Address lines
  // string
  address?: string[];
  
  // ZIP code (if any)
  // string
  zipcode?: string;
  
  // City
  // string
  city?: string;
  
  // State
  // For US, this must be an ISO 3166-2 2-letter state code
  // See https://en.wikipedia.org/wiki/List_of_U.S._state_abbreviations
  // string
  state?: string;
  
  // Country code
  // string
  country_code?: string;
}

// Billing configuration for an organization
export interface BillingConfig {
  // Address of the organization
  // Address
  address?: Address;
  
  // EU VAT number of the organization (if any)
  // string
  vat_number?: string;
  
  // Email address(es) to send emails related to billing (mostly invoices) to.
  // string
  email_addresses?: string[];
  
  // US sales tax number of the organization (if any)
  // string
  us_tax_number?: string;
}

// Request arguments for CreatePaymentMethod
export interface CreatePaymentMethodRequest {
  // The result of PreparePaymentMethod.
  // PreparedPaymentMethod
  prepared_payment_method?: PreparedPaymentMethod;
  
  // First name of owner of payment method
  // string
  first_name?: string;
  
  // Last name of owner of payment method
  // string
  last_name?: string;
}

// An Invoice message describes a transaction for usage of ArangoDB Oasis.
export interface Invoice {
  // System identifier of the invoice.
  // string
  id?: string;
  
  // URL of this resource
  // string
  url?: string;
  
  // Identifier of the organization that is responsible for the payment of this invoice.
  // string
  organization_id?: string;
  
  // Name of the organization that is responsible for the payment of this invoice.
  // string
  organization_name?: string;
  
  // Identifier of the legal entity that is the sender of this invoice.
  // string
  entity_id?: string;
  
  // Name of the legal entity that is the sender of this invoice.
  // string
  entity_name?: string;
  
  // Invoice number (used by accounting)
  // string
  invoice_number?: string;
  
  // The creation date of the invoice
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // If set, this invoice must be manually verified
  // before payment can be initiated.
  // boolean
  requires_manual_verification?: boolean;
  
  // All items of the invoice
  // Invoice_Item
  items?: Invoice_Item[];
  
  // Currency for all amounts
  // string
  currency_id?: string;
  
  // Sum all amount for all items
  // number
  total_amount_excl_vat?: number;
  
  // VAT amount for all items
  // number
  total_vat?: number;
  
  // Sum of total_amount_ex_vat + total_vat.
  // This is the amount that the customer will be charged for.
  // number
  total_amount_incl_vat?: number;
  
  // If set, the VAT reverse charge rule is applied for this invoice.
  // boolean
  vat_reverse_charge?: boolean;
  
  // Invoice_Status
  status?: Invoice_Status;
  
  // All payment attempts for this invoice, ordered by created_at.
  // Invoice_Payment
  payments?: Invoice_Payment[];
}

// A single item of the invoice
export interface Invoice_Item {
  // Identifiers of the UsageItems that this item covers.
  // string
  usageitem_ids?: string[];
  
  // Amount of money (ex VAT) for this item
  // number
  amount?: number;
  
  // Human readable description of this item
  // string
  description?: string;
}

// Payment (attempt) of the invoice
export interface Invoice_Payment {
  // The timestamp of the start of the payment attempt.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Identifier of the payment provider that is used for this payment.
  // string
  payment_provider_id?: string;
  
  // Identifier of this payment (created by payment provider)
  // string
  payment_id?: string;
  
  // Identifier of the payment method that is used for this payment.
  // string
  payment_method_id?: string;
  
  // If set, this payment is still being processed.
  // boolean
  is_pending?: boolean;
  
  // If set, this payment has been payed for succesfully.
  // boolean
  is_completed?: boolean;
  
  // If set, this payment has been rejected.
  // boolean
  is_rejected?: boolean;
  
  // The timestamp of succesfull completion of the payment.
  // googleTypes.Timestamp
  completed_at?: googleTypes.Timestamp;
  
  // The timestamp of rejected completion of the payment.
  // googleTypes.Timestamp
  rejected_at?: googleTypes.Timestamp;
  
  // Human readable reason for the rejection.
  // string
  rejection_reason?: string;
}

// Status of the invoice
export interface Invoice_Status {
  // If set, this invoice is still being processed.
  // boolean
  is_pending?: boolean;
  
  // If set, a successful payment has been made for this invoice.
  // boolean
  is_completed?: boolean;
  
  // If set, all payment attempts for this invoice have been rejected.
  // boolean
  is_rejected?: boolean;
  
  // If set, this invoice has been verified manually.
  // boolean
  is_verified?: boolean;
  
  // The timestamp of succesfull completion of the payment.
  // This field equals the completed_at field of the last payment if
  // that payment succeeded, nil otherwise.
  // googleTypes.Timestamp
  completed_at?: googleTypes.Timestamp;
  
  // The timestamp of rejected completion of the payment.
  // This field equals the rejected_at field of the last payment if
  // that payment failed, nil otherwise.
  // googleTypes.Timestamp
  rejected_at?: googleTypes.Timestamp;
}

// List of Invoices.
export interface InvoiceList {
  // Invoice
  items?: Invoice[];
}

// Request arguments for ListInvoices
export interface ListInvoicesRequest {
  // Request invoices for the organization with this id.
  // This is a required field.
  // string
  organization_id?: string;
  
  // Request invoices that are created at or after this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  from?: googleTypes.Timestamp;
  
  // Request invoices that are created before this timestamp.
  // This is an optional field.
  // googleTypes.Timestamp
  to?: googleTypes.Timestamp;
  
  // Standard list options
  // This is an optional field.
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Request arguments for ListPaymentMethods
export interface ListPaymentMethodsRequest {
  // Identifier of the organization for which payment methods are requested.
  // string
  organization_id?: string;
  
  // Optional common list options. (Context ID is ignored)
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}

// Request arguments for ListPaymentProviders
export interface ListPaymentProvidersRequest {
  // Identifier of the organization for which payment providers are requested.
  // string
  organization_id?: string;
  
  // Optional common list options. (Context ID is ignored)
  // arangodb.cloud.common.v1.ListOptions
  options?: arangodb_cloud_common_v1_ListOptions;
}
export interface PDFDocument {
  // Content of the document
  // string
  content?: string;
  
  // Filename of the document
  // string
  filename?: string;
}

// Payment methods are specific methods for paying at a specific payment provider
// such as a specific credit card.
export interface PaymentMethod {
  // System identifier of this payment method.
  // string
  id?: string;
  
  // Name of the payment method
  // string
  name?: string;
  
  // Description of the payment method
  // string
  description?: string;
  
  // Identifier of the payment provider used for this payment method
  // This is a read-only field.
  // string
  payment_provider_id?: string;
  
  // Identifier of the organization that owns this payment method
  // This is a read-only field.
  // string
  organization_id?: string;
  
  // Creation timestamp of this payment method
  // This is a read-only field.
  // googleTypes.Timestamp
  created_at?: googleTypes.Timestamp;
  
  // Deletion timestamp of this payment method
  // This is a read-only field.
  // googleTypes.Timestamp
  deleted_at?: googleTypes.Timestamp;
  
  // Set if the payment method is deleted.
  // This is a read-only field.
  // boolean
  is_deleted?: boolean;
  
  // If set, this timestamp specifies when the payment method is no longer valid.
  // If not set, there is no (known) end date for this payment method.
  // googleTypes.Timestamp
  valid_until?: googleTypes.Timestamp;
  
  // Token for this payment method, provided by the payment provider.
  // This is a read-only field.
  // string
  token?: string;
  
  // Type of payment method
  // string
  type?: string;
  
  // If set, this payment method is the default for its organization.
  // This is a read-only field.
  // boolean
  is_default?: boolean;
  
  // PaymentMethod_CreditCardInfo
  credit_card_info?: PaymentMethod_CreditCardInfo;
}

// Information of the creditcard.
// Only set when type == "creditcard"
export interface PaymentMethod_CreditCardInfo {
  // Last 4 digits of the CC number.
  // string
  last_digits?: string;
  
  // Type of creditcard
  // string
  card_type?: string;
}

// List of Payment methods
export interface PaymentMethodList {
  // PaymentMethod
  items?: PaymentMethod[];
}

// Payment providers are services that handle payments.
export interface PaymentProvider {
  // System identifier of this payment provider.
  // string
  id?: string;
  
  // Name of the payment provider
  // string
  name?: string;
  
  // Description of the payment provider
  // string
  description?: string;
  
  // Type of payment method supported by this provider
  // string
  type?: string;
}

// List of Payment providers
export interface PaymentProviderList {
  // PaymentProvider
  items?: PaymentProvider[];
}

// Request arguments for PreparePaymentMethod.
export interface PreparePaymentMethodRequest {
  // ID of the provider to prepare
  // string
  provider_id?: string;
  
  // ID of the organization that will own the future payment method
  // string
  organization_id?: string;
}

// Response data for PreparePaymentMethod.
export interface PreparedPaymentMethod {
  // ID of the provider of the future payment method
  // string
  provider_id?: string;
  
  // ID of the organization that will own the future payment method
  // string
  organization_id?: string;
  
  // Token (semantics depends on payment provider)
  // string
  token?: string;
  
  // URL of custom script to load to create the payment method
  // string
  script_url?: string;
  
  // Signature used to verify the consistency of the data in this message.
  // string
  signature?: string;
}

// Request arguments for SetBillingConfig.
export interface SetBillingConfigRequest {
  // Identifier of the organization for which billing address is to be set.
  // string
  organization_id?: string;
  
  // Billing configuration to set.
  // BillingConfig
  config?: BillingConfig;
}

// Request argument for SetDefaultPaymentMethod
export interface SetDefaultPaymentMethodRequest {
  // Identifier of the organization for which the default payment method will be set.
  // string
  organization_id?: string;
  
  // Identifier of the new default payment method for the organization.
  // string
  payment_method_id?: string;
}

// BillingService is the API used to fetch billing information.
export interface IBillingService {
  // Fetch all Invoice resources for the organization identified by the given
  // organization ID that match the given criteria.
  // Required permissions:
  // - billing.invoice.list on the organization identified by the given organization ID
  ListInvoices: (req: ListInvoicesRequest) => Promise<InvoiceList>;
  
  // Fetch a specific Invoice identified by the given ID.
  // Required permissions:
  // - billing.invoice.get on the organization that owns the invoice
  // with given ID.
  GetInvoice: (req: arangodb_cloud_common_v1_IDOptions) => Promise<Invoice>;
  
  // Fetch a specific Invoice identified by the given ID as PDF document.
  // Required permissions:
  // - billing.invoice.get on the organization that owns the invoice
  // with given ID.
  GetInvoicePDF: (req: arangodb_cloud_common_v1_IDOptions) => Promise<PDFDocument>;
  
  // Fetch all payment providers that are usable for the organization identified
  // by the given context ID.
  // Required permissions:
  // - billing.paymentprovider.list on the organization identified by the given context ID
  ListPaymentProviders: (req: ListPaymentProvidersRequest) => Promise<PaymentProviderList>;
  
  // Fetch a specific payment provider identified by the given ID.
  // Required permissions:
  // - None
  GetPaymentProvider: (req: arangodb_cloud_common_v1_IDOptions) => Promise<PaymentProvider>;
  
  // Fetch all payment methods that are configured for the organization identified
  // by the given context ID.
  // Required permissions:
  // - billing.paymentmethod.list on the organization identified by the given context ID
  ListPaymentMethods: (req: ListPaymentMethodsRequest) => Promise<PaymentMethodList>;
  
  // Fetch a specific payment method identified by the given ID.
  // Required permissions:
  // - billing.paymentmethod.get on the organization that owns the payment method
  // which is identified by the given ID
  GetPaymentMethod: (req: arangodb_cloud_common_v1_IDOptions) => Promise<PaymentMethod>;
  
  // Prepare the payment provider for creating a new payment method.
  // Required permissions:
  // - billing.paymentmethod.create on the organization that owns future payment method.
  PreparePaymentMethod: (req: PreparePaymentMethodRequest) => Promise<PreparedPaymentMethod>;
  
  // Create a new payment method.
  // Required permissions:
  // - billing.paymentmethod.create on the organization that owns the given payment method.
  CreatePaymentMethod: (req: CreatePaymentMethodRequest) => Promise<PaymentMethod>;
  
  // Update a specific payment method.
  // Note that only name, description & valid period are updated.
  // Required permissions:
  // - billing.paymentmethod.update on the organization that owns the given payment method.
  UpdatePaymentMethod: (req: PaymentMethod) => Promise<PaymentMethod>;
  
  // Delete a specific payment method identified by the given ID.
  // Required permissions:
  // - billing.paymentmethod.delete on the organization that owns the given payment method
  // which is identified by the given ID.
  DeletePaymentMethod: (req: arangodb_cloud_common_v1_IDOptions) => Promise<void>;
  
  // Fetch the default PaymentMethod for an organization identified by the given ID.
  // Required permissions:
  // - billing.paymentmethod.get-default on the organization that is identified by the given ID
  GetDefaultPaymentMethod: (req: arangodb_cloud_common_v1_IDOptions) => Promise<PaymentMethod>;
  
  // Update the default PaymentMethod for an organization identified by the
  // given organization ID, to the payment method identified by the given payment method ID.
  // Required permissions:
  // - billing.paymentmethod.set-default on the organization identified by the given organization ID
  SetDefaultPaymentMethod: (req: SetDefaultPaymentMethodRequest) => Promise<void>;
  
  // Fetch the billing configuration of an organization identified by the given ID.
  // Required permissions:
  // - billing.config.get on the organization that is identified by the given ID
  GetBillingConfig: (req: arangodb_cloud_common_v1_IDOptions) => Promise<BillingConfig>;
  
  // Update the billing configuration for an organization identified by the
  // given organization ID.
  // Required permissions:
  // - billing.config.set on the organization identified by the given organization ID
  SetBillingConfig: (req: SetBillingConfigRequest) => Promise<void>;
}

// BillingService is the API used to fetch billing information.
export class BillingService implements IBillingService {
  // Fetch all Invoice resources for the organization identified by the given
  // organization ID that match the given criteria.
  // Required permissions:
  // - billing.invoice.list on the organization identified by the given organization ID
  async ListInvoices(req: ListInvoicesRequest): Promise<InvoiceList> {
    const path = `/api/billing/v1/organization/${encodeURIComponent(req.organization_id || '')}/invoices`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific Invoice identified by the given ID.
  // Required permissions:
  // - billing.invoice.get on the organization that owns the invoice
  // with given ID.
  async GetInvoice(req: arangodb_cloud_common_v1_IDOptions): Promise<Invoice> {
    const path = `/api/billing/v1/invoices/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific Invoice identified by the given ID as PDF document.
  // Required permissions:
  // - billing.invoice.get on the organization that owns the invoice
  // with given ID.
  async GetInvoicePDF(req: arangodb_cloud_common_v1_IDOptions): Promise<PDFDocument> {
    const path = `/api/billing/v1/invoices/${encodeURIComponent(req.id || '')}/pdf`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all payment providers that are usable for the organization identified
  // by the given context ID.
  // Required permissions:
  // - billing.paymentprovider.list on the organization identified by the given context ID
  async ListPaymentProviders(req: ListPaymentProvidersRequest): Promise<PaymentProviderList> {
    const path = `/api/billing/v1/organization/${encodeURIComponent(req.organization_id || '')}/paymentproviders`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific payment provider identified by the given ID.
  // Required permissions:
  // - None
  async GetPaymentProvider(req: arangodb_cloud_common_v1_IDOptions): Promise<PaymentProvider> {
    const path = `/api/billing/v1/paymentproviders/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Fetch all payment methods that are configured for the organization identified
  // by the given context ID.
  // Required permissions:
  // - billing.paymentmethod.list on the organization identified by the given context ID
  async ListPaymentMethods(req: ListPaymentMethodsRequest): Promise<PaymentMethodList> {
    const path = `/api/billing/v1/organization/${encodeURIComponent(req.organization_id || '')}/paymentmethods`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
  
  // Fetch a specific payment method identified by the given ID.
  // Required permissions:
  // - billing.paymentmethod.get on the organization that owns the payment method
  // which is identified by the given ID
  async GetPaymentMethod(req: arangodb_cloud_common_v1_IDOptions): Promise<PaymentMethod> {
    const path = `/api/billing/v1/paymentmethods/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Prepare the payment provider for creating a new payment method.
  // Required permissions:
  // - billing.paymentmethod.create on the organization that owns future payment method.
  async PreparePaymentMethod(req: PreparePaymentMethodRequest): Promise<PreparedPaymentMethod> {
    const url = `/api/billing/v1/paymentproviders/${encodeURIComponent(req.provider_id || '')}/prepare`;
    return api.post(url, req);
  }
  
  // Create a new payment method.
  // Required permissions:
  // - billing.paymentmethod.create on the organization that owns the given payment method.
  async CreatePaymentMethod(req: CreatePaymentMethodRequest): Promise<PaymentMethod> {
    const url = `/api/billing/v1/paymentmethods`;
    return api.post(url, req);
  }
  
  // Update a specific payment method.
  // Note that only name, description & valid period are updated.
  // Required permissions:
  // - billing.paymentmethod.update on the organization that owns the given payment method.
  async UpdatePaymentMethod(req: PaymentMethod): Promise<PaymentMethod> {
    const url = `/api/billing/v1/paymentmethods/${encodeURIComponent(req.id || '')}`;
    return api.put(url, req);
  }
  
  // Delete a specific payment method identified by the given ID.
  // Required permissions:
  // - billing.paymentmethod.delete on the organization that owns the given payment method
  // which is identified by the given ID.
  async DeletePaymentMethod(req: arangodb_cloud_common_v1_IDOptions): Promise<void> {
    const path = `/api/billing/v1/paymentmethods/${encodeURIComponent(req.id || '')}`;
    const url = path + api.queryString(req, [`id`]);
    return api.delete(url, undefined);
  }
  
  // Fetch the default PaymentMethod for an organization identified by the given ID.
  // Required permissions:
  // - billing.paymentmethod.get-default on the organization that is identified by the given ID
  async GetDefaultPaymentMethod(req: arangodb_cloud_common_v1_IDOptions): Promise<PaymentMethod> {
    const path = `/api/billing/v1/organization/${encodeURIComponent(req.id || '')}/default-paymentmethod`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Update the default PaymentMethod for an organization identified by the
  // given organization ID, to the payment method identified by the given payment method ID.
  // Required permissions:
  // - billing.paymentmethod.set-default on the organization identified by the given organization ID
  async SetDefaultPaymentMethod(req: SetDefaultPaymentMethodRequest): Promise<void> {
    const url = `/api/billing/v1/organization/${encodeURIComponent(req.organization_id || '')}/default-paymentmethod`;
    return api.put(url, req);
  }
  
  // Fetch the billing configuration of an organization identified by the given ID.
  // Required permissions:
  // - billing.config.get on the organization that is identified by the given ID
  async GetBillingConfig(req: arangodb_cloud_common_v1_IDOptions): Promise<BillingConfig> {
    const path = `/api/billing/v1/organization/${encodeURIComponent(req.id || '')}/config`;
    const url = path + api.queryString(req, [`id`]);
    return api.get(url, undefined);
  }
  
  // Update the billing configuration for an organization identified by the
  // given organization ID.
  // Required permissions:
  // - billing.config.set on the organization identified by the given organization ID
  async SetBillingConfig(req: SetBillingConfigRequest): Promise<void> {
    const url = `/api/billing/v1/organization/${encodeURIComponent(req.organization_id || '')}/config`;
    return api.put(url, req);
  }
}
