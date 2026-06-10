.PHONY: build run-server clean 

build: clean
	go fmt ./...
	go vet ./...
	go build -o cmd/server/server ./cmd/server
	go build -o cmd/shorten/shorten ./cmd/shorten

run-server: build  
	./cmd/server/server

clean:
	rm -f cmd/server/server
	rm -f cmd/shorten/shorten
	
