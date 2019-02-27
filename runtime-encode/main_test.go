package main

import "testing"

func TestEncode(t *testing.T) {
	tests := []struct {
		input  string
		output string
		error  string
	}{
		{"aabbbcadddd", "2a3b1c1a4d", ""},
		{"aaaaaaaaaa", "10a", ""},
		{"ab", "1a1b", ""},
		{"aa", "2a", ""},
		{"a", "1a", ""},
		{"__+", "2_1+", ""},
		{"", "", ""},
		{"1a", "", "String cannot contain numbers"},
	}

	for _, test := range tests {
		result, err := Encode(test.input)
		if result != test.output {
			t.Errorf("%s should have returned %s", test.input, test.output)
		}
		if err != nil && err.Error() != test.error {
			t.Errorf("Expected error: %s", test.error)
		}
	}
}
