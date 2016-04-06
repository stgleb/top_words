package main

import (
	"flag"
	"sync"
	"top-words/top_words"
)

var addr = flag.String("addr", ":8000", "http service address")
var port = flag.String("port", "9000", "tcp service port")
var host = flag.String("host", "localhost", "tcp service host")


func main() {
	var wg sync.WaitGroup
	// Run tcp server
	wg.Add(2)
	go top_words.RunTCPServer(*port, *host, &wg)
	go top_words.RunHTTPServer(*addr, &wg)
	wg.Wait()
}