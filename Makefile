output = ./serve
ifeq ($(GOOS),windows)
  output = './serve.exe'
endif

all:
	go run ./bin

build:
	go build -o $(output) ./bin

spelling:
	pyspelling
