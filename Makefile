export PATH := $(PWD)/build:$(PATH)

run: build run-tmp-main

run-tmp-main:
	muto run tmp/main.mu

build: build-cli

build-cli:
	go build -o build/muto ./cmd/cli
