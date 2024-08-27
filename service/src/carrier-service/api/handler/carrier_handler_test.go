package handler

import (
	"testing"
)

func TestCarrierHandler_CreateCarrier(t *testing.T) {
	tests := map[string]struct {
		description  string // description of the test case
		route        string // route path to test
		expectedCode int
	}{
		"empty name":      {"", "test", true, "carrier name is empty"},
		"empty address":   {"test", "", true, "carrier address is empty"},
		"active is false": {"test", "test", false, "carrier is not active"},
	}
}
