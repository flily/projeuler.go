package framework

import (
	"testing"

	"github.com/fatih/color"
)

func TestStyleOnInteger(t *testing.T) {
	cases := []struct {
		style    Style
		input    any
		expected string
	}{
		{
			style:    NewIntegerStyle(0, 0),
			input:    42,
			expected: "42",
		},
		{
			style:    NewIntegerStyle(5, 0),
			input:    42,
			expected: "   42",
		},
		{
			style:    NewIntegerStyle(5, 0).WithPadding("0"),
			input:    42,
			expected: "00042",
		},
		{
			style:    NewIntegerStyle(5, 0).Left(),
			input:    42,
			expected: "42   ",
		},
		{
			style:    NewIntegerStyle(5, 0).Center(),
			input:    42,
			expected: " 42  ",
		},
		{
			style:    NewIntegerStyle(6, 0).Center(),
			input:    42,
			expected: "  42  ",
		},
		{
			style:    NewIntegerStyle(0, 0),
			input:    -42,
			expected: "-42",
		},
		{
			style:    NewIntegerStyle(5, 0),
			input:    -42,
			expected: "  -42",
		},
	}

	for _, c := range cases {
		result := c.style.Apply(c.input)
		if result != c.expected {
			t.Errorf("wrong result for '%v'", c.input)
			t.Errorf("expect: %s", c.expected)
			t.Errorf("got   : %s", result)
		}
	}
}

func TestStyleOnFloat(t *testing.T) {
	cases := []struct {
		style    Style
		input    any
		expected string
	}{
		{
			style:    NewFloatStyle(0, 2),
			input:    3.1415926,
			expected: "3.14",
		},
		{
			style:    NewFloatStyle(6, -1),
			input:    3.1415926,
			expected: "3.141593",
		},
		{
			style:    NewFloatStyle(6, 2),
			input:    3.1415926,
			expected: "  3.14",
		},
		{
			style:    NewFloatStyle(6, 2).WithPadding("0"),
			input:    3.1415926,
			expected: "003.14",
		},
	}

	for _, c := range cases {
		result := c.style.Apply(c.input)
		if result != c.expected {
			t.Errorf("wrong result for '%v'", c.input)
			t.Errorf("expect: %s", c.expected)
			t.Errorf("got   : %s", result)
		}
	}
}

func TestStyleOnGeneric(t *testing.T) {
	cases := []struct {
		style    Style
		input    any
		expected string
	}{
		{
			style:    NewGenericStyle(0),
			input:    "lorem",
			expected: "lorem",
		},
		{
			style:    NewGenericStyle(10),
			input:    "lorem",
			expected: "lorem     ",
		},
		{
			style:    NewGenericStyle(10).WithPadding("0"),
			input:    "lorem",
			expected: "lorem00000",
		},
		{
			style:    NewGenericStyle(10).Center(),
			input:    "lorem",
			expected: "  lorem   ",
		},
		{
			style:    NewGenericStyle(11).Center(),
			input:    "lorem",
			expected: "   lorem   ",
		},
		{
			style:    NewGenericStyle(10).Right(),
			input:    "lorem",
			expected: "     lorem",
		},
		{
			style:    NewGenericStyle(10),
			input:    42,
			expected: "42        ",
		},
		{
			style:    NewGenericStyle(10).Center(),
			input:    42,
			expected: "    42    ",
		},
		{
			style:    NewGenericStyle(10).Right(),
			input:    42,
			expected: "        42",
		},
		{
			style:    NewGenericStyle(10),
			input:    []int{1, 2, 3},
			expected: "[1 2 3]   ",
		},
		{
			style:    NewGenericStyle(10).Center(),
			input:    []int{1, 2, 3},
			expected: " [1 2 3]  ",
		},
		{
			style:    NewGenericStyle(10).Right(),
			input:    []int{1, 2, 3},
			expected: "   [1 2 3]",
		},
	}

	for _, c := range cases {
		result := c.style.Apply(c.input)
		if result != c.expected {
			t.Errorf("wrong result for '%v'", c.input)
			t.Errorf("expect: %s", c.expected)
			t.Errorf("got   : %s", result)
		}
	}
}

func TestColorLibrary(t *testing.T) {
	c := color.New(color.FgRed, color.BgGreen)
	c.EnableColor()
	got := c.Sprintf("lorem")
	exp := "\x1b[31;42mlorem\x1b[0;0m"
	if got != exp {
		t.Errorf("wrong result")
		t.Errorf("expect: %s", exp)
		t.Errorf("got   : %s", got)
	}
}

func TestStyleColourString(t *testing.T) {
	cases := []struct {
		style    Style
		input    any
		expected string
	}{
		{
			style:    NewGenericStyle(10).Colour(ColourRed).Force(),
			input:    "error",
			expected: "\x1b[31;0merror     \x1b[0;0m",
		},
	}

	for _, c := range cases {
		result := c.style.Apply(c.input)
		if result != c.expected {
			t.Errorf("wrong result for '%v'", c.input)
			t.Errorf("expect: %s", c.expected)
			t.Errorf("got   : %s", result)
		}
	}
}
