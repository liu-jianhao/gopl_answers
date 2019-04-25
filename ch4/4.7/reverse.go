package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("usage: %s str\n")
		return
	}
	b := []byte(os.Args[1])
	reverse(b)
	fmt.Printf("%s\n", b)
}

func reverse(s []byte) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
