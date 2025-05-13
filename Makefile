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
