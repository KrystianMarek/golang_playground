all: clean client server

.PHONY: client
client:
	echo client
	cd client && go build -o ../bin/client

.PHONY: server
server:
	echo server
	cd server && go build -o ../bin/server

.PHONY: clean
clean:
	rm -rf bin/*

