#!/bin/bash
# Use this script to install all needed dependencies for ubuntu

apt install -y git
apt install -y make
apt install -y curl
apt install -y jq
GO111MODULE=on go get github.com/golangci/golangci-lint/cmd/golangci-lint@v1.14.1
go get -v -u github.com/client9/misspell/cmd/misspell
go get -v -u github.com/fzipp/gocyclo
go get -v -u github.com/golang/dep/cmd/dep
go get -v -u golang.org/x/lint/golint
go get -v -u github.com/gordonklaus/ineffassign
go get -v -u github.com/h12w/gosweep
go get -v -u github.com/mattn/goveralls
go get -v -u github.com/stripe/safesql
go get -v -u golang.org/x/crypto/ssh/terminal
go get -v -u golang.org/x/net/html
go get -v -u golang.org/x/text
go get -v -u golang.org/x/tools/cmd/goimports
go get -v -u github.com/ory/go-acc
apt install -y udate-ca-certificates
