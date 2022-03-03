package encoder

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetMaxStr(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		n        int
		expected string
	}{
		{
			name:     "emptyN",
			str:      "",
			n:        10,
			expected: "",
		},
		{
			name:     "empty0",
			str:      "",
			n:        0,
			expected: "",
		},
		{
			name:     "stringSmaller",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        10,
			expected: "abcdefghij",
		},
		{
			name:     "stringLarger",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        50,
			expected: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name:     "stringEqual",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        26,
			expected: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name:     "accentsSmaller",
			str:      "áéíóúÁÚÍÓÚñÑ",
			n:        5,
			expected: "áéíóú",
		},
		{
			name:     "accentsLarger",
			str:      "áéíóúÁÚÍÓÚñÑ",
			n:        50,
			expected: "áéíóúÁÚÍÓÚñÑ",
		},
		{
			name:     "accentsEqual",
			str:      "áéíóúÁÚÍÓÚñÑ",
			n:        12,
			expected: "áéíóúÁÚÍÓÚñÑ",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getMaxStr(test.str, test.n)
			if actual != test.expected {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}

func TestGetMaxStrEllipsis(t *testing.T) {
	tests := []struct {
		name     string
		str      string
		n        int
		expected string
	}{
		{
			name:     "emptyN",
			str:      "",
			n:        10,
			expected: "",
		},
		{
			name:     "empty0",
			str:      "",
			n:        0,
			expected: "",
		},
		{
			name:     "stringSmaller",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        10,
			expected: "abcdefg...",
		},
		{
			name:     "stringLarger",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        50,
			expected: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name:     "stringEqual",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        26,
			expected: "abcdefghijklmnopqrstuvwxyz",
		},
		{
			name:     "accentsSmaller",
			str:      "áéíóúÁÚÍÓÚñÑ",
			n:        8,
			expected: "áéíóú...",
		},
		{
			name:     "accentsLarger",
			str:      "áéíóúÁÚÍÓÚñÑ",
			n:        50,
			expected: "áéíóúÁÚÍÓÚñÑ",
		},
		{
			name:     "accentsEqual",
			str:      "áéíóúÁÚÍÓÚñÑ",
			n:        12,
			expected: "áéíóúÁÚÍÓÚñÑ",
		},
		{
			name:     "NoEllipsis0",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        0,
			expected: "",
		},
		{
			name:     "NoEllipsis1",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        1,
			expected: "a",
		},
		{
			name:     "NoEllipsis2",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        2,
			expected: "ab",
		},
		{
			name:     "NoEllipsis3",
			str:      "abcdefghijklmnopqrstuvwxyz",
			n:        3,
			expected: "abc",
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual := getMaxStrEllipsis(test.str, test.n)
			if actual != test.expected {
				assert.Equal(t, test.expected, actual)
			}
		})
	}
}
