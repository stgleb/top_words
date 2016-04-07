# Top words web service


## Installation:

### Download sources

1.1 git clone https://github.com/stgleb/top_words.git

### Install dependencies

1.2 make install



## Basic usage:

2.1 ###Build server

    make

2.2 ### Run server

    ./main


2.3 ### Push words to service via tcp protocol:

	echo "go bla bla-bla bla foo foo foo bar boo" | nc localhost 9000

2.4 ### Request for n most frequest words:

	curl http://localhost:8000/?N=<int>



## Running tests:

    cd top_words
    go test -v
