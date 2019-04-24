package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s filename\n", os.Args[0])
		return
	}

	wordfreq(os.Args[1])
}

func wordfreq(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println("open file error")
		return
	}
	counts := make(map[string]int)
	input := bufio.NewScanner(file)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}

	fmt.Printf("world \t\t counts \n")
	for k, v := range counts {
		fmt.Printf("%s \t\t %d\n", k, v)
	}
}
