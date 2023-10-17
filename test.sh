#! /bin/bash 

go test ./... -cover -coverprofile=coverage.out -covermode=atomic
