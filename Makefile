.PHONY: run build

APP_FILE=cmd/main.go


run:
	go run $(APP_FILE)

build:
	go build -o  /usr/bin/rb_keygen $(APP_FILE)
