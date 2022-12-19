.PHONY: all clean windows linux wasm arm test

all: linux windows wasm

linux: main.go
	GOOS=linux go build -o Obstacle

windows: main.go
	GOOS=windows go build -o Obstacle.exe

wasm: main.go web.go
	GOOS=js GOARCH=wasm go build -o Obstacle.wasm -ldflags "-X main.iointerface=js"
	cp Obstacle.wasm www/
	cp "`go env GOROOT`/misc/wasm/wasm_exec.js" www/

arm: main.go
	CC=arm-linux-gnueabi-gcc CGO_ENABLED=1 GOOS=linux GOARM=7 GOARCH=arm go build -o Obstacle-arm

test: vintbas
	go test -run=TestMain

clean:
	rm -f Obstacle Obstacle.exe Obstacle.wasm Obstacle-arm
