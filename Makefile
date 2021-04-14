PACKAGE_NAME := "ghq.alfredworkflow"
EXEC_BIN := "alfred-ghq"
DIST_DIR := "build"
.PHONY: all mod copy-build-assets package-workflow clean

all: build copy-build-assets package-workflow

mod:
	go mod download

build: mod
	CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -o $(DIST_DIR)/$(EXEC_BIN)-amd64
	CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -o $(DIST_DIR)/$(EXEC_BIN)-arm64

copy-build-assets:
	cp info.plist icon.png $(DIST_DIR)

package-workflow:
	cd $(DIST_DIR) && zip -r $(PACKAGE_NAME) ./

clean:
	rm -r $(DIST_DIR)
