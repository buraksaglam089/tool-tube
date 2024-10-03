build: 
	@go build -o bin/tool-tube

run: build
	@./bin/tool-tube

test:
	@go test -v ./...