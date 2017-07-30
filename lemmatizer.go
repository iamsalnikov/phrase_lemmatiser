package main

import (
	"strings"
	"unicode/utf8"

	"github.com/dveselov/mystem"
)

type Lemmatizer struct{}

func (l *Lemmatizer) normalizeString(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

func (l *Lemmatizer) getWords(str string) []string {
	words := strings.Split(str, " ")
	var result []string

	for _, word := range words {
		word = strings.TrimSpace(word)
		if utf8.RuneCountInString(word) > 0 {
			result = append(result, word)
		}
	}

	return result
}

func (l *Lemmatizer) GetLemmas(str string) []string {
	str = l.normalizeString(str)
	words := l.getWords(str)

	lemmasMap := make(map[string]bool)
	for _, word := range words {
		analyses := mystem.NewAnalyses(word)

		for i := 0; i < analyses.Count(); i++ {
			lemma := analyses.GetLemma(i)
			lemmasMap[lemma.Text()] = true
		}

		analyses.Close()
	}

	lemmas := make([]string, 0, len(lemmasMap))
	for l, _ := range lemmasMap {
		lemmas = append(lemmas, l)
	}

	return lemmas
}
