package collections

import (
	"testing"
)

func TestStringSet_Has(t *testing.T) {
	testTable := []map[string]map[string]bool{
		{
			"src": {},
			"expected": {
				"hello": false,
			},
		},

		{
			"src": {
				"hello": true,
				"world": true,
			},
			"expected": {
				"hello": true,
				"gmail": false,
				"world": true,
			},
		},
	}

	for _, testCase := range testTable {
		set := NewStringSet()
		for w := range testCase["src"] {
			set.Add(w)
		}

		for w, exists := range testCase["expected"] {
			if set.Has(w) != exists {
				t.Errorf("We expected that set %v will contain string \"%s\"", set.GetData(), w)
			}
		}
	}
}

func TestStringSet_GetData(t *testing.T) {
	testTable := []map[string][]string{
		{
			"src":      {},
			"expected": {},
		},
		{
			"src":      {"hello", "world"},
			"expected": {"hello", "world"},
		},
		{
			"src":      {"hello", "world", "world", "gmail", "hello"},
			"expected": {"hello", "gmail", "world"},
		},
	}

	for _, testCase := range testTable {
		set := NewStringSet()

		for _, w := range testCase["src"] {
			set.Add(w)
		}

		if !IsStringSlicesEquals(set.GetData(), testCase["expected"]) {
			t.Errorf("Slices are not identical. Expected %v, got %v", testCase["expected"], set.GetData())
		}
	}
}
