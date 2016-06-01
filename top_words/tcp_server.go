package top_words

import (
	"log"
	"net"
	"os"
	"sync"
)

const TCP = "tcp"

func RunTCPServer(port string, host string, wg *sync.WaitGroup) {
	// Listen for incoming connections.
	l, err := net.Listen(TCP, host+":"+port)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Println("Listening on " + host + ":" + TCP)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		log.Println("Accepted connection from", conn.RemoteAddr())

		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn)
	}
	wg.Done()
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	// Read the incoming connection into the buffer.
	buffer := make([]byte, 0, 4096)
	log.Println("Handle connection")

	for {
		// Big buffer for all data.
		// Small buffer for reading portions of data.
		tmp := make([]byte, 1024)
		reqLen, err := conn.Read(tmp)
		log.Println("Read ", reqLen, " bytes")

		if err != nil {
			log.Println("Error reading:", err.Error())
			break
		}
		buffer = append(buffer, tmp[:reqLen]...)
		log.Println("Data received ", buffer)
	}

	words := ParseString(buffer)
	log.Println("Words ", words)

	for i := 0; i < len(words); i++ {
		shard := wordsMap.GetShard(words[i])
		shard.Aux.Lock()
		log.Println("Critical section begin")
		count, ok := wordsMap.Get(words[i])
		cnt := 0

		if ok {
			cnt = count.(int)
			cnt = cnt + 1
		}
		log.Println(cnt)
		wordsMap.Set(words[i], cnt)

		shard.Aux.Unlock()
		log.Println("Critical section end")
	}

	log.Println("Finish connection handling")
	defer conn.Close()
}
