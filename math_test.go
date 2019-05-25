package selection

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBytesWithinTolerance(t *testing.T) {
	tests := []struct {
		a         byte
		b         byte
		tolerance byte
		result    bool
	}{
		{a: 0, b: 0, tolerance: 0, result: true},
		{a: 0, b: 0, tolerance: 1, result: true},
		{a: 0, b: 255, tolerance: 255, result: true},
		{a: 255, b: 0, tolerance: 255, result: true},
		{a: 255, b: 255, tolerance: 0, result: true},
		{a: 255, b: 255, tolerance: 1, result: true},

		{a: 0, b: 255, tolerance: 0, result: false},
		{a: 0, b: 255, tolerance: 254, result: false},
		{a: 255, b: 0, tolerance: 0, result: false},
		{a: 255, b: 0, tolerance: 254, result: false},
	}

	for _, test := range tests {
		result := BytesWithinTolerance(test.a, test.b, test.tolerance)
		assert.Equal(t, test.result, result)
	}
}
