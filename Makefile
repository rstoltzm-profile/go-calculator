BINARY=go-calculator

all: build

build: build-linux build-windows

build-linux:
	CG0_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o bin/$(BINARY)-linux main.go

build-windows:
	GOOS=windows GOARCH=amd64 go build -o bin/$(BINARY)-windows.exe main.go

test: 
	go test -v

clean:
	rm -f bin/$(BINARY)-windows.exe main.go bin/$(BINARY)-linux main.go

test-bin: build-linux
	./test-go-calculator-cli.sh