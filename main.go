package top_words


import (
    "github.com/gorilla/mux"
    "flag"
    "net/http"
    "log"
)

var addr = flag.String("addr", ":9000", "http service address")
var port = flag.String("port", "8000", "tcp service port")
var host = flag.String("host", "localhost", "tcp service host")


func main() {
	// Run tcp server
	resultChannel := make(chan string)
	Serve(port, host, resultChannel)

	// Run http server
	r := mux.NewRouter()
    r.HandleFunc("/", Get(HomeHandler))

    if err := http.ListenAndServe(*addr, nil); err != nil {
        log.Fatal("ListenAndServe:", err)
    }
}