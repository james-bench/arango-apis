//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

const (
	// PaymentMethod types
	// Used in PaymentProvider.type & PaymentMethod.type.

	// PaymentMethodTypeCreditCard indicates that the payment method
	// is of type creditcard.
	PaymentMethodTypeCreditCard = "creditcard"
)

const (
	// Credit card types
	// Used in PaymentMethod.CreditCardInfo.card_type;

	// CardTypeVisa indicates a VISA creditcard
	CardTypeVisa = "visa"
	// CardTypeMastercard indicates a Mastercard creditcard
	CardTypeMastercard = "mastercard"
)
