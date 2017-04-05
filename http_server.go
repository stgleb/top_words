package top_words

import (
	"bytes"
	"context"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"sync"
	"time"
)

var server *http.Server

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	N := r.URL.Query().Get("N")
	logger.Printf("Got request from %s", r.RemoteAddr)

	if N == "" {
		N = "0"
	}
	logger.Println("N parameter ", N)
	count, err := strconv.Atoi(N)

	if err != nil {
		http.Error(w, "Wrong number format", http.StatusUnprocessableEntity)
	}

	words := TopN(count)
	logger.Println(words)

	var buffer bytes.Buffer
	logger.Println("Words response ", words)

	for i := 0; i < len(words); i++ {
		buffer.WriteString(words[i])
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(buffer.Bytes())
}

func RunHTTPServer(addr string, wg *sync.WaitGroup) {
	// Run http server
	logger.Println("Start http server on ", addr)
	router := mux.NewRouter()

	router.HandleFunc("/", HomeHandler).Methods("GET")
	http.Handle("/", router)

	server = &http.Server{Addr: ":8081", Handler: router}

	go func() {
		// service connections
		if err := server.ListenAndServe(); err != nil {
			logger.Fatal("ListenAndServe:", err)
		}
	}()

	wg.Done()
}

func ShutDownHTTP() {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	if server != nil {
		server.Shutdown(ctx)
	}
}
