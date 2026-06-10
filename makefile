.PHONY: build 

build:
		go fmt ./...
		go vet ./...
		go build -o cmd/server/server ./cmd/server
		go build -o cmd/shorten/shorten ./cmd/shorten

run server: 
		cmd/server/server

clean:
	rm cmd/server/server
	rm cmd/shorten/shorten
