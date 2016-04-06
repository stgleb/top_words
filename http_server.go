package top_words

import (
	"net/http"
	"github.com/gorilla/context"
	"github.com/streamrail/concurrent-map"
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

	wordsMap := context.Get(r, WORDS_MAP)
	words := TopN(N, wordsMap)

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(words)
}


// Decorator that accepts only get  methods
func Get(h http.Handler) http.Handler {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "GET" {
			http.Error(w, "Method not allowed", 405)
			return
		} else {
			h.ServeHTTP(w, r)
		}
	}

	return wrapper
}


func CountMapContext(h http.Handler, wordsMap *cmap.ConcurrentMap) http.Handler {
	wrapper := func(w http.ResponseWriter, r *http.Request) {
		_, err := context.GetOk(r, WORDS_MAP)

		if err != false {
			context.Set(r, WORDS_MAP, wordsMap)
		}

		h.ServeHTTP(w, r)

	}
	return wrapper
}