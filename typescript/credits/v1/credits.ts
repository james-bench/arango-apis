//
// This file is AUTO-GENERATED by protoc-gen-ts.
// Do not modify it manually.
///
import api from '../../api'
import * as googleTypes from '../../googleTypes'

// File: credits/v1/credits.proto
// Package: arangodb.cloud.credits.v1
export interface CreditBundle {
  // Unique identifier of this credit bundle.
  // string
  id?: string;
  
  // The organization this credit bundle belongs to.
  // string
  organization_id?: string;
  
  // The number of credits purchased in this bundle.
  // number
  credits_purchased?: number;
  
  // The total price of these credits.
  // number
  total_price?: number;
  
  // Currency used for total_price.
  // string
  currency?: string;
  
  // The number of credits remaining in this bundle.
  // number
  credits_remaining?: number;
  
  // The date at which this bundle was purchased.
  // googleTypes.Timestamp
  purchased_at?: googleTypes.Timestamp;
  
  // The date from which this bundle is valid.
  // googleTypes.Timestamp
  valid_from?: googleTypes.Timestamp;
  
  // The date until which this bundle is valid.
  // googleTypes.Timestamp
  valid_until?: googleTypes.Timestamp;
}

// List of credit bundles
export interface CreditBundlesList {
  // CreditBundle
  items?: CreditBundle[];
}

// Request for listing credit bundles
export interface ListCreditBundlesRequest {
  // ID of the organization for which credit bundles are listed.
  // This is a required field.
  // string
  organization_id?: string;
  
  // If set, include expired bundles.
  // boolean
  include_expired?: boolean;
}

// CreditsService is the API used for managing credits.
export interface ICreditsService {
  // List credit bundles for an organization.
  // Required permissions:
  // - bundle.credits.list on the organization identified by the given organization ID
  ListCreditBundles: (req: ListCreditBundlesRequest) => Promise<CreditBundlesList>;
}

// CreditsService is the API used for managing credits.
export class CreditsService implements ICreditsService {
  // List credit bundles for an organization.
  // Required permissions:
  // - bundle.credits.list on the organization identified by the given organization ID
  async ListCreditBundles(req: ListCreditBundlesRequest): Promise<CreditBundlesList> {
    const path = `/api/credits/v1/${encodeURIComponent(req.organization_id || '')}/creditbundles`;
    const url = path + api.queryString(req, [`organization_id`]);
    return api.get(url, undefined);
  }
}
