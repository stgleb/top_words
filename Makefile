all:
	go build main.go
install:
	go get github.com/gorilla/mux
	go get github.com/stgleb/concurrent-map
	go get github.com/stretchr/testify/assert 
clean:
	rm main
