Top words web service

Basic usage:

1) Push words to service via tcp protocol:

	echo "go bla bla-bla bla foo foo foo bar boo" | nc localhost 9000

2) Request for n most frequest words:

	curl http://localhost:8000/?N=<int>



Running tests:

    cd top_words
    go test -v
