.PHONY: build
build: clean proto
	go build -o bin/protobuf

.PHONY: proto
proto:
	protoc proto/echo/echo.proto --go_out=./ --go-json_out=./

.PHONY: clean
clean:
	rm -rf proto/echo/*.go
	rm -rf bin/*

