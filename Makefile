output = ./course
ifeq ($(GOOS),windows)
  output = './course.exe'
endif

all: build
build:
	go build -o $(output) ./src/

spelling:
	pyspelling
