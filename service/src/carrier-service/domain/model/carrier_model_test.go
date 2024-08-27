package model

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestCarrier_Successful(t *testing.T) {
	t.Run("create new carrier", func(t *testing.T) {
		actual := Carrier{}
		err := actual.NewCarrier("name", "address", true)
		require.NoError(t, err)

		assert.Equal(t, "name", actual.Name)
		assert.Equal(t, "address", actual.Address)
		assert.Equal(t, true, actual.Active)
		assert.False(t, actual.CreatedAt.IsZero())
		assert.False(t, actual.UpdatedAt.IsZero())
	})

	t.Run("update carrier address", func(t *testing.T) {
		actual := Carrier{}
		err := actual.NewCarrier("name", "address", true)
		require.NoError(t, err)
		assert.Equal(t, "address", actual.Address)

		err = actual.UpdateCarrierAddress("test")
		require.NoError(t, err)
		assert.Equal(t, "test", actual.Address)

	})

	t.Run("update carrier active", func(t *testing.T) {
		actual := Carrier{}
		err := actual.NewCarrier("name", "address", true)
		require.NoError(t, err)
		assert.True(t, actual.Active)

		actual.UpdateCarrierActiveStatus(false)
		assert.False(t, actual.Active)

	})

}

func TestCarrier_Error(t *testing.T) {
	t.Run("create new carrier fail", func(t *testing.T) {
		tests := map[string]struct {
			name          string
			address       string
			active        bool
			expectedError string
		}{
			"empty name":      {"", "test", true, "carrier name is empty"},
			"empty address":   {"test", "", true, "carrier address is empty"},
			"active is false": {"test", "test", false, "carrier is not active"},
		}

		for _, tc := range tests {
			got := Carrier{}
			err := got.NewCarrier(tc.name, tc.address, tc.active)
			require.EqualError(t, err, tc.expectedError)
		}

	})

	t.Run("update carrier address fail", func(t *testing.T) {
		actual := Carrier{}
		err := actual.NewCarrier("name", "address", true)
		require.NoError(t, err)
		assert.Equal(t, "address", actual.Address)

		err = actual.UpdateCarrierAddress("")
		require.Error(t, err, "carrier address is empty")
		assert.Equal(t, "address", actual.Address)
	})
}
