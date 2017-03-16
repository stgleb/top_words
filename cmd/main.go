package main

import (
	"flag"
	"fmt"
	"github.com/stgleb/top_words"
	"log"
	"net/http"
	_ "net/http/pprof"
	"sync"
)

var addr = flag.String("addr", ":8000", "http service address")
var port = flag.String("port", "9000", "tcp service port")
var host = flag.String("host", "0.0.0.0", "tcp service host")
var pprofEnabled = flag.Bool("pprof", false, "Enable pprof server")
var pprofPort = flag.Int("pprofPort", 8080, "Pprof http server port")

func runPprof() {
	addr := fmt.Sprintf("0.0.0.0:%d", *pprofPort)
	log.Printf("Starting pprof server on http://%s", addr)
	log.Println(http.ListenAndServe(addr, nil))
}

func init() {
	flag.Parse()
}

func main() {
	var wg sync.WaitGroup
	// Run tcp server
	wg.Add(2)
	// Start pprof HTTP server.
	if *pprofEnabled {
		go runPprof()
	}
	go top_words.RunTCPServer(*port, *host, &wg)
	go top_words.RunHTTPServer(*addr, &wg)
	wg.Wait()
}
