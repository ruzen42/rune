BUILD_DIR=bin
VERSION=0.1.0
NAME=rune
PREFIX=/usr/local

.PHONY: all build clean run fmt install

all: build

install: build 
	mkdir -p $(PREFIX)/bin
	chmod u+s $(BUILD_DIR)/$(NAME)
	cp $(BUILD_DIR)/$(NAME) $(PREFIX)/bin/$(NAME)

build: fmt
	@mkdir -p $(BUILD_DIR)
	go build -ldflags "-X main.Version=$(VERSION)" -o $(BUILD_DIR)/$(NAME) .  
	sudo setcap cap_sys_admin,cap_setuid,cap_setgid+ep $(BUILD_DIR)/$(NAME)

clean:
	rm -rf $(BUILD_DIR)

run: build
	./$(BUILD_DIR)/$(NAME)

fmt:
	go fmt
