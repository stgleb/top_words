package main

import (
	"flag"
	"fmt"
	"github.com/stgleb/top_words"
	"log"
	"net/http"
	_ "net/http/pprof"
	"os"
	"os/signal"
	"sync"
)

var (
	addr         = flag.String("addr", ":8000", "http service address")
	port         = flag.String("port", "9000", "tcp service port")
	host         = flag.String("host", "0.0.0.0", "tcp service host")
	pprofEnabled = flag.Bool("pprof", false, "Enable pprof server")
	pprofPort    = flag.Int("pprofPort", 8080, "Pprof http server port")
)

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

	sigInt := make(chan os.Signal, 1)
	signal.Notify(sigInt, os.Interrupt)

	go func() {
		for range sigInt {
			log.Println("Shutdown on SIGINT")
			top_words.ShutDownHTTP()
			top_words.ShutDownTCP()
			// sig is a ^C, handle it
		}
	}()

	wg.Wait()
}
