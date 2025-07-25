.PHONY: run build go-run

APP_NAME = gomi.exe
DIST_DIR = dist

build:
	mkdir -p $(DIST_DIR)
	go build -ldflags="-H=windowsgui" -o $(DIST_DIR)/$(APP_NAME) .

run-buid: build
	./$(DIST_DIR)/$(APP_NAME)

run:
	go run main.go
