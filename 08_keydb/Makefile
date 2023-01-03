MyKeyDB = "MyKeyDB"

.PHONY: build
build: clean
	go build -o bin/keydb

.PHONY: keydb-start
keydb-start: keydb-stop
	docker run -d --name ${MyKeyDB} -v $$(pwd)/db/data:/data -v $$(pwd)/db/keydb.conf:/etc/keydb/keydb.conf -p 6379:6379 eqalpha/keydb

.PHONY: keydb-stop
keydb-stop:
	docker rm -f ${MyKeyDB} || true

.PHONY: clean
clean:
	rm -rf bin/*

