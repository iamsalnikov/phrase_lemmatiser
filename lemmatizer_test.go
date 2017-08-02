package main

import (
	"github.com/iamsalnikov/phrase_lemmatiser/collections"
	"testing"
)

func TestLemmatizer_normalizeString(t *testing.T) {
	lemmatizer := Lemmatizer{}
	testTable := map[string]string{"": "", "Hello World": "hello world"}

	for src, expected := range testTable {
		if expected != lemmatizer.normalizeString(src) {
			t.Errorf("Expect \"%s\" but got \"%s\"", expected, src)
		}
	}
}

func TestLemmatizer_getWords(t *testing.T) {
	lemmatizer := Lemmatizer{}
	testTable := []map[string]interface{}{
		{
			"src":      "",
			"expected": []string{},
		},
		{
			"src":      "hello",
			"expected": []string{"hello"},
		},
		{
			"src":      "hello     world email world hello email",
			"expected": []string{"hello", "world", "email"},
		},
	}

	for _, testCase := range testTable {
		src, ok := testCase["src"].(string)
		if !ok {
			t.Error("Can not get string from src")
			continue
		}

		expected, ok := testCase["expected"].([]string)
		if !ok {
			t.Error("Can not get []string from expected")
		}

		words := lemmatizer.getWords(src)
		if !collections.IsStringSlicesEquals(words, expected) {
			t.Errorf("We expected that %v and %v will be equals", words, expected)
		}
	}
}
