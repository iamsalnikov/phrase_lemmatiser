package main

import (
	"strings"
	"unicode/utf8"

	"github.com/dveselov/mystem"
	"github.com/iamsalnikov/phrase_lemmatiser/collections"
)

type Lemmatizer struct{}

func (l *Lemmatizer) normalizeString(str string) string {
	return strings.TrimSpace(strings.ToLower(str))
}

func (l *Lemmatizer) getWords(str string) []string {
	words := strings.Split(str, " ")
	set := collections.NewStringSet()

	for _, word := range words {
		word = strings.TrimSpace(word)
		if utf8.RuneCountInString(word) > 0 {
			set.Add(word)
		}
	}

	return set.GetData()
}

func (l *Lemmatizer) GetLemmas(str string) []string {
	str = l.normalizeString(str)
	words := l.getWords(str)

	lemmas := collections.NewStringSet()
	for _, word := range words {
		analyses := mystem.NewAnalyses(word)

		for i := 0; i < analyses.Count(); i++ {
			lemma := analyses.GetLemma(i)
			lemmas.Add(lemma.Text())
		}

		analyses.Close()
	}

	return lemmas.GetData()
}
