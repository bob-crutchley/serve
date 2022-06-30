build:
	go build serve.go

install: build
	sudo cp serve /usr/local/bin

uninstall:
	sudo rm -f /usr/local/bin/serve
