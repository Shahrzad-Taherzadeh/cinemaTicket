BUILD_DIRECTORY := build

build_directory:
	mkdir -p $(BUILD_DIRECTORY)

linux: build_directory
	GOOS=linux GOARCH=amd64 go build -o $(BUILD_DIRECTORY)/wallet-manager-linux-amd64 cmd/main.go

darwin: build_directory
	GOOS=darwin GOARCH=amd64 go build -o $(BUILD_DIRECTORY)/wallet-manager-darwin-amd64 cmd/main.go

windows: build_directory
	GOOS=windows GOARCH=amd64 go build -o $(BUILD_DIRECTORY)/wallet-manager-windows-amd64.exe cmd/main.go

all: linux darwin windows

api-doc:
# 	swag init -g cmd/main.go -o internal/api/docs
	~/Documents/go_projects/bin/swag init -g cmd/main.go