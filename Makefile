export PATH := $(PWD)/build:$(PATH)

run: build run-tmp-main

repl: build repl-tmp-main


run-explain: build run-tmp-main-explain

run-tmp-main-explain:
	muto run --explain tmp/main.mu

repl-tmp-main:
	muto repl

run-tmp-main:
	muto run tmp/main.mu

example-tictactoe: build
	muto run examples/tictactoe.mu

build: build-cli

build-cli:
	go build -o build/muto ./cmd/cli

install-latest:
	go install github.com/SSripilaipong/muto@$(shell git rev-parse origin/main)

go-get-common:
	go get github.com/SSripilaipong/go-common@$(shell curl -s https://api.github.com/repos/SSripilaipong/go-common/commits/main | jq -r .sha)
