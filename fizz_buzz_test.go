package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSimpleText_Text(t *testing.T) {
	tests := []struct {
		text string
	}{
		{"Fizz"},
		{"Buzz"},
		{"FizzBuzz"},
	}
	for _, test := range tests {
		s := &SimpleText{text: test.text}
		assert.Equalf(t, test.text, s.Text(), "input %s", test.text)
	}
}

func TestValue_String(t *testing.T) {
	tests := []struct {
		name     string
		value    Value
		expected string
	}{
		{
			name: "FizzBuzz is nil -> original",
			value: Value{
				Value:    nil,
				Original: 20,
			},
			expected: "20",
		},
		{
			name: "FizzBuzzValue = {Fizz} -> Fizz",
			value: Value{
				Value:    &SimpleText{text: "Fizz"},
				Original: 12,
			},
			expected: "Fizz",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := test.value.String()
			if actual != test.expected {
				t.Errorf("expected = %v, actual = %v", test.expected, actual)
			}
		})
	}
}

func TestSingleFailure(t *testing.T) {
	value := Value{
		Value:    &SimpleText{"Foo"},
		Original: 3,
	}
	assert.Equal(t, "Fizz", value.String())
}
