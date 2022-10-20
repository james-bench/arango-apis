//
// DISCLAIMER
//
// Copyright 2022 ArangoDB GmbH, Cologne, Germany
//

package v1

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	t.Run("NewMessage", func(t *testing.T) {
		m := NewMessage("test reason", "test dberror", "test syncerror")
		assert.Equal(t, "test reason", m.Reason)
		assert.Equal(t, "test dberror", m.DbErrors)
		assert.Equal(t, "test syncerror", m.SyncErrors)
		t.Run("AsJSON", func(t *testing.T) {
			str := `{"reason":"test reason","db_errors":"test dberror","sync_errors":"test syncerror"}`
			assert.Equal(t, str, m.AsJSON())
		})
	})
	t.Run("NewMessageFromString", func(t *testing.T) {
		var m Message
		err := m.FromJSON(`
	    {"reason":"test reason",
	    "db_errors":"test dberror",
	    "sync_errors":"test syncerror"}
	    `)
		assert.Nil(t, err)
		assert.Equal(t, "test reason", m.Reason)
		assert.Equal(t, "test dberror", m.DbErrors)
		assert.Equal(t, "test syncerror", m.SyncErrors)

		t.Run("invalid", func(t *testing.T) {
			var m Message
			// missing ','
			err := m.FromJSON(`
            {"reason":"test reason",
            "db_errors":"test dberror"
            "sync_errors":"test syncerror"}
            `)
			assert.NotNil(t, err)
		})
	})
}
