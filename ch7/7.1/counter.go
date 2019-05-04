package main

import (
	"bufio"
	"fmt"
	"strings"
)

type WordCounter int

func (c *WordCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(strings.NewReader(string(p)))
	scan.Split(bufio.ScanWords)

	for scan.Scan() {
		(*c)++
	}
	return len(p), nil
}

type LineCounter int

func (c *LineCounter) Write(p []byte) (int, error) {
	scan := bufio.NewScanner(strings.NewReader(string(p)))
	scan.Split(bufio.ScanLines)

	for scan.Scan() {
		(*c)++
	}
	return len(p), nil
}

func main() {
	var wc WordCounter
	wc.Write([]byte("hello world golang")) // 3
	fmt.Println(wc)

	var lc LineCounter
	lc.Write([]byte("hello\nworld\ngolang")) // 3
	fmt.Println(lc)
}
