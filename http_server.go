package top_words

import (
	"net/http"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	N := r.URL.Query().Get("N")

	if N == "" {
		N = 0
	}

	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return

	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write("response")
}


// Decorator that accepts only get  methods
func Get(next func(w http.ResponseWriter, r *http.Request)) func(w http.ResponseWriter, r *http.Request){
	wrapper := func (w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		} else {
			return next(w, r)
		}
	}

	return wrapper
}