test:
	go test -v ./...

clean:
	rm -rf dist/
	
dist/bin:
	mkdir -p dist/bin

build: dist/bin
	go build ./... -o dist/bin/
