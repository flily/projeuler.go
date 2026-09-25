package management

import (
	"testing"
)

func TestEntryName(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello world", "HelloWorld"},
		{"foo_bar", "FooBar"},
		{"lorem-ipsum", "LoremIpsum"},
		{"123abc", "123abc"},
	}

	for _, test := range tests {
		result := MakeEntryName(test.input)
		if result != test.expected {
			t.Errorf("MakeEntryName(%q) = %q; want %q", test.input, result, test.expected)
		}
	}
}
