package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestUnpacking(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{{
		input:    "a4",
		expected: "aaaa",
	},
		{
			input:    "ba4",
			expected: "baaaa",
		},
		{
			input:    "a" + string(rune(92)) + "4",
			expected: "a4",
		},
		{
			input:    "aa",
			expected: "aa",
		},
		{
			input:    string(rune(92)) + "44",
			expected: "4444",
		},
		{
			input:    "aa",
			expected: "aa",
		},
		{
			input:    "aa4",
			expected: "aaaaa",
		},
		{
			input:    "a10",
			expected: "aaaaaaaaaa",
		},
		{
			input:    string(rune(92)) + string(rune(92)) + "4",
			expected: string(rune(92)) + string(rune(92)) + string(rune(92)) + string(rune(92)),
		},

		{
			input:    "qwe" + string(rune(92)) + "4" + string(rune(92)) + "5",
			expected: "qwe45",
		},
		{
			input:    "qwe" + string(rune(92)) + "45",
			expected: "qwe44444",
		},
		{
			input:    "qwe" + string(rune(92)) + string(rune(92)) + "5",
			expected: "qwe" + string(rune(92)) + string(rune(92)) + string(rune(92)) + string(rune(92)) + string(rune(92)),
		},
	}
	for _, i := range tests {
		actual, err := unpacking(i.input)
		assert.Nil(t, err)
		assert.Equal(t, i.expected, actual)
	}
}
