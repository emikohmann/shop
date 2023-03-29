package util

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestIsEmpty(t *testing.T) {
	for _, testCase := range []struct {
		name     string
		input    func() interface{}
		expected bool
	}{
		{
			name: "Nil input",
			input: func() interface{} {
				return nil
			},
			expected: true,
		},
		{
			name: "Empty map",
			input: func() interface{} {
				return map[string]interface{}{}
			},
			expected: true,
		},
		{
			name: "Non-empty map",
			input: func() interface{} {
				return map[string]interface{}{
					"test": "value",
				}
			},
			expected: false,
		},
		{
			name: "Empty slice",
			input: func() interface{} {
				return []interface{}{}
			},
			expected: true,
		},
		{
			name: "Non-empty slice",
			input: func() interface{} {
				return []interface{}{
					"test", "data",
				}
			},
			expected: false,
		},
		{
			name: "Empty channel",
			input: func() interface{} {
				return make(chan interface{})
			},
			expected: true,
		},
		{
			name: "Non-empty channel",
			input: func() interface{} {
				ch := make(chan interface{}, 1)
				ch <- "test data"
				return ch
			},
			expected: false,
		},
		{
			name: "Empty pointer",
			input: func() interface{} {
				var input *string = nil
				return input
			},
			expected: true,
		},
		{
			name: "Non-empty pointer",
			input: func() interface{} {
				inputString := "test data"
				return &inputString
			},
			expected: false,
		},
		{
			name: "Empty variable",
			input: func() interface{} {
				return ""
			},
			expected: true,
		},
		{
			name: "Non-empty variable",
			input: func() interface{} {
				return "test data"
			},
			expected: false,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			assert.Equal(t, testCase.expected, IsEmpty(testCase.input()))
		})
	}
}
