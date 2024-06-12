output = ./course
ifeq ($(GOOS),windows)
  output = './course.exe'
endif

all: build
build:
	go build -o $(output) ./src/

generate:
	./course generate ./public

spelling:
	pyspelling

zip:
	rm -rf /tmp/digging_deeper/
	./course generate /tmp/digging_deeper/
	cd /tmp/ && zip -r course.zip  digging_deeper/
