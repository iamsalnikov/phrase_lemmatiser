package collections

import "testing"

func TestIsStringSlicesEquals(t *testing.T) {
	testTable := []map[string]interface{}{
		{
			"a":      []string{},
			"b":      []string{},
			"equals": true,
		},
		{
			"a":      []string{},
			"b":      []string{"hello"},
			"equals": false,
		},
		{
			"a":      []string{"hello", "world"},
			"b":      []string{"world", "hello"},
			"equals": true,
		},
		{
			"a":      []string{"hello", "hello"},
			"b":      []string{"hello"},
			"equals": false,
		},
		{
			"a":      []string{"hello", "hello"},
			"b":      []string{"hello", "hello"},
			"equals": true,
		},
	}

	for _, testCase := range testTable {
		a, aok := testCase["a"].([]string)
		b, bok := testCase["b"].([]string)
		e, eok := testCase["equals"].(bool)
		if !aok || !bok || !eok {
			t.Error("Can not type cast")
			continue
		}

		if IsStringSlicesEquals(a, b) != e {
			result := "will be equals"
			if !e {
				result = "will not be equals"
			}
			t.Errorf("We expected that slices %v and %v "+result, testCase["a"].([]string), testCase["b"].([]string))
		}
	}
}
