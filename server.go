package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Server struct{}

func (s *Server) Run(addr string) {
	lemmatizer := Lemmatizer{}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		phrase := r.URL.Query().Get("phrase")
		data := lemmatizer.GetLemmas(phrase)

		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		answer, _ := json.Marshal(data)
		fmt.Fprintln(w, string(answer))
	})

	http.ListenAndServe(addr, nil)
}
