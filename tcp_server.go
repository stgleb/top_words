package top_words

import (
	"net"
	"os"
	"strings"
	"log"
)

const TCP  = "tcp"


func Serve(port string, host string,resultChannel chan string) {
	// Listen for incoming connections.
	l, err := net.Listen(TCP, host + ":" + port)
	if err != nil {
		log.Fatal("Error listening:", err.Error())
	}
	// Close the listener when the application closes.
	defer l.Close()
	log.Println("Listening on " + host + ":" + TCP)
	for {
		// Listen for an incoming connection.
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting: ", err.Error())
			os.Exit(1)
		}
		// Handle connections in a new goroutine.
		go handleRequest(conn, resultChannel)
	}
}

// Handles incoming requests.
func handleRequest(conn net.Conn,resultChannel chan string) {
	// Make a buffer to hold incoming data.
	// Read the incoming connection into the buffer.
	buffer := make([]byte, 0, 4096)

	for {
		// Big buffer for all data.
		// Small buffer for reading portions of data.
		tmp := make([]byte, 1024)
		reqLen, err := conn.Read(tmp)

		if err != nil {
			log.Println("Error reading:", err.Error())
			break
		}

		if reqLen == 0 {
			break
		}

		buffer = append(buffer, tmp[:reqLen]...)
	}

	words := parseString(buffer)

	for i := 0; i < len(words); i++ {
		resultChannel <- words[i]
	}

	defer conn.Close()
}

func parseString(bytes []byte) []string {
	s := string(bytes)

	return strings.SplitAfter(s, ' ')
}