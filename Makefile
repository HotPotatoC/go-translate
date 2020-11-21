MAIN_SRC = main.go
OUTFILE = bin/translate

.PHONY: help clean build build-all

# Source: https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## Displays all the available commands
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

clean: ## Deletes all compiled / executable files
	@find bin -type f -name 'translate*' -print0 | xargs -0 rm --
	@echo ">> Deleted all build files!"

build: ## Compile the go files
	@echo ">> Building go file..."
	@go build -ldflags="-s -w" -o $(OUTFILE) $(MAIN_SRC)

build-all: ## Compile the go files for multiple OS
	@echo ">> Building go files for multiple OS..."
	GOOS=linux GOARCH=arm go build -ldflags="-s -w" -o $(OUTFILE)-linux-arm $(MAIN_SRC)
	GOOS=linux GOARCH=arm64 go build -ldflags="-s -w" -o $(OUTFILE)-linux-arm64 $(MAIN_SRC)
	GOOS=linux GOARCH=386 go build -ldflags="-s -w" -o $(OUTFILE)-linux-386 $(MAIN_SRC)
	GOOS=freebsd GOARCH=386 go build -ldflags="-s -w" -o $(OUTFILE)-freebsd-386 $(MAIN_SRC)
	GOOS=windows GOARCH=386 go build -ldflags="-s -w" -o $(OUTFILE)-windows-386 $(MAIN_SRC)
	echo ">> Finished building"
	@ls -hl -d ./$(OUTFILE)* ./bin/

.DEFAULT_GOAL := help