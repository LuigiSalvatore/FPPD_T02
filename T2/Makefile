.PHONY: all build
BUILD_PATH = 	./Builds/
S_PATH = 		./server/
C_PATH = 		./client/

all: build

go.mod:
	go mod init jogo
	go get -u github.com/nsf/termbox-go

build: go.mod
	go build $(S_PATH)jogo_s.go
	go build $(C_PATH)jogo_c.go
	
clean:
	rm -f jogo

distclean: clean
	rm -f go.mod go.sum

run:
	make 
	jogo.exe