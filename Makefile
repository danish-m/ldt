.PHONY: all
all: build

APP=ldt
APP_EXECUTABLE="./out/$(APP)"

clean: ##@clean remove executable
	rm -rf out/

compile:
	mkdir -p out/
	GO111MODULE=on go build -o $(APP_EXECUTABLE)

build: clean compile ##@build a fresh build

linux-build: clean
	GOOS=linux GOARCH=amd64 go build -o $(APP_EXECUTABLE)
