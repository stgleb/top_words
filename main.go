package top_words

import (
	"flag"
	"net/http"
	"log"
	"github.com/gorilla/mux"
	"github.com/streamrail/concurrent-map"
)

var addr = flag.String("addr", ":9000", "http service address")
var port = flag.String("port", "8000", "tcp service port")
var host = flag.String("host", "localhost", "tcp service host")


func main() {
	wordsMap := cmap.New()
	// Run tcp server
	Serve(port, host, wordsMap)

	// Run http server
	r := mux.NewRouter()
	handler := CountMapContext(Get(HomeHandler), wordsMap)
	r.HandleFunc("/", handler())

	if err := http.ListenAndServe(*addr, nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}