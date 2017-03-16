package top_words

import (
	"bytes"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
	"sync"
)

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	N := r.URL.Query().Get("N")
	log.Printf("Got request from %s", r.RemoteAddr)

	if N == "" {
		N = "0"
	}
	log.Println("N parameter ", N)
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

func RunHTTPServer(addr string, wg *sync.WaitGroup) {
	// Run http server
	log.Println("Start http server on ", addr)
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	http.Handle("/", router)

	srv := &http.Server{Addr: ":8081", Handler: router}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	// TODO(stgleb): Add SIGHUP handling for graceful shutdown
	// srv.Shutdown(context.Background())
	wg.Done()
}