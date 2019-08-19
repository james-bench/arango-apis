//
// DISCLAIMER
//
// Copyright 2019 ArangoDB Inc, Cologne, Germany
//
// Author Ewout Prangsma
//

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInvoiceURL(t *testing.T) {
	assert.Equal(t, "/Organization/123/Invoice/u1", InvoiceURL("123", "u1"))
	assert.Equal(t, "/Organization/123%2F567/Invoice/c2%2F3", InvoiceURL("123/567", "c2/3"))
	assert.Equal(t, "/Organization/123%2F567/Invoice/a%25b", InvoiceURL("123/567", "a%b"))
}
