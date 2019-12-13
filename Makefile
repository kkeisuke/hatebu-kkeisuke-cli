COMMAND := htb

.PHONY: build
build:
	go build -o bin/$(COMMAND)
