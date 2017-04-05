package top_words

import (
	"log"
	"net"
	"os"
	"sync"
)

const TCP = "tcp"

var (
	shutdown chan struct{}
	logger   *log.Logger
)

func init() {
	shutdown = make(chan struct{})
	logger = log.New(os.Stdout,
		"INFO: ",
		log.Ldate|log.Ltime|log.Lshortfile)
}

func ShutDownTCP() {
	close(shutdown)
}

func RunTCPServer(port string, host string, wg *sync.WaitGroup) {
	// Listen for incoming connections.
	l, err := net.Listen(TCP, host+":"+port)
	if err != nil {
		logger.Fatal("Error listening:", err.Error())
	}
	// Close the listener when the application closes.
	defer l.Close()
	logger.Println("Listening on " + TCP + "://" + host + ":" + port)
	for {
		select {
		case <-shutdown:
			return
		default:
			// Listen for an incoming connection.
			conn, err := l.Accept()
			logger.Println("Accepted connection from", conn.RemoteAddr())

			if err != nil {
				logger.Fatalf("Error accepting: %v", err)
			}
			// Handle connections in a new goroutine.
			go handleRequest(conn)
		}
	}
	wg.Done()
}

// Handles incoming requests.
func handleRequest(conn net.Conn) {
	// Make a buffer to hold incoming data.
	// Read the incoming connection into the buffer.
	buffer := make([]byte, 0, 4096)
	logger.Println("Handle connection")

	for {
		// Big buffer for all data.
		// Small buffer for reading portions of data.
		tmp := make([]byte, 1024)
		reqLen, err := conn.Read(tmp)
		logger.Println("Read ", reqLen, " bytes")

		if err != nil {
			logger.Println("Error reading:", err.Error())
			break
		}
		buffer = append(buffer, tmp[:reqLen]...)
		logger.Println("Data received ", buffer)
	}

	words := ParseString(buffer)
	logger.Println("Words ", words)

	for i := 0; i < len(words); i++ {
		shard := wordsMap.GetShard(words[i])
		shard.Aux.Lock()
		logger.Println("Critical section begin")
		count, ok := wordsMap.Get(words[i])
		cnt := 0

		if ok {
			cnt = count.(int)
			cnt = cnt + 1
		}
		logger.Println(cnt)
		wordsMap.Set(words[i], cnt)

		shard.Aux.Unlock()
		logger.Println("Critical section end")
	}

	logger.Println("Finish connection handling")
	defer conn.Close()
}
