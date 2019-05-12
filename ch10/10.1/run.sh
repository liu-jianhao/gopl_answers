#!/usr/bin/env bash

go build ./mandelbrot.go
go build ./jpeg.go

./mandelbrot | ./jpeg -output=gif > test.gif

