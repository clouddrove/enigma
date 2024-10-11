package docker

import (
	"testing"
)

func TestIsValidEnvVarKey(t *testing.T) {
	testCases := []struct {
		name     string
		key      string
		expected bool
	}{
		{
			name:     "Empty string",
			key:      "",
			expected: false,
		},
		{
			name:     "Single letter",
			key:      "A",
			expected: true,
		},
		{
			name:     "Single underscore",
			key:      "_",
			expected: true,
		},
		{
			name:     "Letter and digit",
			key:      "MY_VAR1",
			expected: true,
		},
		{
			name:     "Letter, digit, and underscore",
			key:      "MY_VAR_123",
			expected: true,
		},
		{
			name:     "Starting with digit",
			key:      "123_MY_VAR",
			expected: false,
		},
		{
			name:     "Invalid character",
			key:      "MY_VAR@",
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isValidEnvVarKey(tc.key)
			if result != tc.expected {
				t.Errorf("Expected %v for key %q, but got %v", tc.expected, tc.key, result)
			}
		})
	}
}

func TestIsLetter(t *testing.T) {
	testCases := []struct {
		name     string
		r        rune
		expected bool
	}{
		{
			name:     "Uppercase letter",
			r:        'A',
			expected: true,
		},
		{
			name:     "Lowercase letter",
			r:        'a',
			expected: true,
		},
		{
			name:     "Underscore",
			r:        '_',
			expected: true,
		},
		{
			name:     "Digit",
			r:        '1',
			expected: false,
		},
		{
			name:     "Special character",
			r:        '@',
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isLetter(tc.r)
			if result != tc.expected {
				t.Errorf("Expected %v for rune %q, but got %v", tc.expected, string(tc.r), result)
			}
		})
	}
}

func TestIsLetterOrDigitOrUnderscore(t *testing.T) {
	testCases := []struct {
		name     string
		r        rune
		expected bool
	}{
		{
			name:     "Uppercase letter",
			r:        'A',
			expected: true,
		},
		{
			name:     "Lowercase letter",
			r:        'a',
			expected: true,
		},
		{
			name:     "Underscore",
			r:        '_',
			expected: true,
		},
		{
			name:     "Digit",
			r:        '1',
			expected: true,
		},
		{
			name:     "Special character",
			r:        '@',
			expected: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := isLetterOrDigitOrUnderscore(tc.r)
			if result != tc.expected {
				t.Errorf("Expected %v for rune %q, but got %v", tc.expected, string(tc.r), result)
			}
		})
	}
}
