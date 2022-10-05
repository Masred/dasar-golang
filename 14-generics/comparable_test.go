package golang_generics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func IsSame[T comparable](value1, value2 T) bool {
	if value1 == value2 {
		return true
	} else {
		return false
	}
}

func TestIsSame(t *testing.T) {
	assert.True(t, IsSame("masred", "masred"))
	assert.True(t, IsSame(100, 100))
	assert.True(t, IsSame(true, true))
}
