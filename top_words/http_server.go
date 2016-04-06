package top_words

import (
	"net/http"
	"strconv"
	"bytes"
	"log"
	"sync"
	"github.com/gorilla/mux"
)


func HomeHandler(w http.ResponseWriter, r *http.Request) {
	N := r.URL.Query().Get("N")

	if N == "" {
		N = "0"
	}
	log.Println("N parameter ",N)
	count, err := strconv.Atoi(N)

	if err != nil {
		http.Error(w, "Wrong number format", 422)
	}

	words := TopN(count)
	log.Println(words)

	var buffer bytes.Buffer
	log.Println("Words response ", words)

    for i := 0; i < len(words); i++ {
        buffer.WriteString(words[i])
    }

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buffer.Bytes())
}


func RunHTTPServer(addr string, wg *sync.WaitGroup){
	// Run http server
	log.Println("Start http server on ",addr)
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	http.Handle("/", router)

	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
	wg.Done()
}