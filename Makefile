.PHONY: setup build

build:
	go build -o agones-factorio-sdk cmd/main.go

setup:
	dep ensure -vendor-only
