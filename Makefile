output = ./course
ifeq ($(GOOS),windows)
  output = './course.exe'
endif

all: build
build:
	go build -o $(output) ./scripts/

spelling:
	pyspelling
