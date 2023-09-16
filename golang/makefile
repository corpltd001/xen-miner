.PHONY: build*

dev:
	go run .

build-all: build-windows-amd64 \
	build-windows-arm64 \
	build-linux-amd64 \
	build-linux-arm64 \
	build-macos-amd64 \
	build-darwin-arm64 \
	build-local

build-local:
	go build -o build/xenminer .

build: build-local

build-macos-amd64:
	GOOS=darwin GOARCH=amd64 go build -o build/xenminer-macos-amd64 .

build-darwin-arm64:
	GOOS=darwin GOARCH=arm64 go build -o build/xenminer-macos-arm64 .

build-windows-amd64:
	GOOS=windows GOARCH=amd64 go build -o build/xenminer-windows-amd64.exe .

build-windows-arm64:
	GOOS=windows GOARCH=arm64 go build -o build/xenminer-windows-arm64.exe .

build-linux-amd64:
	GOOS=windows GOARCH=amd64 go build -o build/xenminer-linux-amd64 .

build-linux-arm64:
	GOOS=linux GOARCH=arm64 go build -o build/xenminer-linux-arm64 .

test:
	go test -v --count 1 .

bench:
	go test --bench .
