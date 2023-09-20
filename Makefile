SHELL 			:= /bin/bash
VERSION 		:= 1.0
KIND_CLUSTER    := subhrajit-starter-cluster
KIND            := kindest/node:v1.27.3


run:
	go run main.go

test:
	go test ./... -count=1
	staticcheck -checks=all ./...

all:
