package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

// 'пятак', 'пятка' и 'тяпка' - принадлежат одному множеству,
// 'листок', 'слиток' и 'столик' - другому.
func TestAnagrams(t *testing.T) {
	tests := []struct {
		input    []string
		expected map[string][]string
	}{{
		input:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик"},
		expected: map[string][]string{"пятак": []string{"пятка", "тяпка"}, "листок": []string{"слиток", "столик"}},
	},
		{
			input:    []string{"Пятак", "пЯтка", "тяПка", "лИстОк", "слиТок", "Столик"},
			expected: map[string][]string{"пятак": []string{"пятка", "тяпка"}, "листок": []string{"слиток", "столик"}},
		},
		{
			input:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "столик", "ф"},
			expected: map[string][]string{"пятак": []string{"пятка", "тяпка"}, "листок": []string{"слиток", "столик"}},
		},
		{
			input:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "слиток"},
			expected: map[string][]string{"пятак": []string{"пятка", "тяпка"}, "листок": []string{"слиток"}},
		},
		{
			input:    []string{"пятак", "пятка", "тяпка", "листок", "слиток", "слиток"},
			expected: map[string][]string{"пятак": []string{"пятка", "тяпка"}, "листок": []string{"слиток"}},
		},
	}
	for _, i := range tests {
		actual := Anagrams(i.input)
		assert.Equal(t, i.expected, actual)
	}
}
