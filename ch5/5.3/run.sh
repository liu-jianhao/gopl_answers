#!/usr/bin/env bash

go build ../../ch1/1.8/fetch.go
go build ./findlinks.go
# ./fetch golang.org | ./findlinks
./fetch www.baidu.com | ./findlinks
