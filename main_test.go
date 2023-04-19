package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLargestElementInSlice(t *testing.T) {
	tests := map[string]struct {
		input    []int32
		expected int32
	}{
		"expect largest number to be returned": {
			input:    []int32{3, 5, 9, 2, 1, 49},
			expected: 49,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			actual := LargestElementInSlice(tt.input)
			assert.Equal(t, tt.expected, actual)
		})
	}
}
