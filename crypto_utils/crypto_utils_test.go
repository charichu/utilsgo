package crypto_utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAccessTokenConstants(t *testing.T) {
	if expirationTime != 24 {
		t.Error("expiration time should be 24 hours")
	}
	assert.EqualValues(t, 24, expirationTime, "expiration time should be 24 hours")
}
